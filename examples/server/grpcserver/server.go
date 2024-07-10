package grpcserver

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/geniusrabbit/udetect/examples/server/middleware"
	"github.com/geniusrabbit/udetect/examples/server/tools"
	"github.com/geniusrabbit/udetect/protocol"
)

type contextWrapper func(context.Context) context.Context

// GRPCServer wrapper object
type GRPCServer struct {
	Detector       protocol.DetectorServer
	RequestTimeout time.Duration
	ContextWrap    contextWrapper
	Logger         *zap.Logger
	// Secure connection certificates
	CertFile string
	KeyFile  string
}

// RunGRPC server
func (s *GRPCServer) RunGRPC(ctx context.Context, listen string) error {
	network, address := parseNetwork(listen)
	s.Logger.Info("Start GRPC API: " + network + " " + address)

	l, err := net.Listen(network, address)
	if err != nil {
		return err
	}

	zapOpts := []grpc_zap.Option{
		grpc_zap.WithDurationField(func(duration time.Duration) zapcore.Field {
			return zap.Int64("grpc.time_ns", duration.Nanoseconds())
		}),
	}

	// Init certification
	creds, err := loadCreds(s.CertFile, s.KeyFile)
	if err != nil {
		if closeErr := l.Close(); err != nil {
			s.Logger.Error("failed to close",
				zap.String(`network`, network), zap.String(`address`, address), zap.Error(closeErr))
		}
		return errors.Wrap(err, `failed to setup TLS:`)
	}

	srv := grpc.NewServer(
		// grpc.ConnectionTimeout(s.RequestTimeout),
		grpc.Creds(creds),
		grpc_middleware.WithUnaryServerChain(
			grpc_zap.UnaryServerInterceptor(s.Logger, zapOpts...),
			middleware.GRPCContextUnaryWrapper(s.contextWrapFnk()),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_zap.StreamServerInterceptor(s.Logger, zapOpts...),
			middleware.GRPCContextStreamWrapper(s.contextWrapFnk()),
		),
	)
	protocol.RegisterDetectorServer(srv, s.Detector)

	go func() {
		<-ctx.Done()
		srv.GracefulStop()
		if closeErr := l.Close(); err != nil {
			s.Logger.Error("failed to close",
				zap.String(`network`, network), zap.String(`address`, address), zap.Error(closeErr))
		}
	}()

	s.Logger.Info(fmt.Sprintf("Starting listening at %s", listen))
	return srv.Serve(l)
}

// RunHTTP server
func (s *GRPCServer) RunHTTP(ctx context.Context, address string) error {
	s.Logger.Info("Start HTTP API: " + address)

	// @link https://grpc-ecosystem.github.io/grpc-gateway/docs/mapping/customizing_your_gateway/
	gw := runtime.NewServeMux(
		runtime.WithMarshalerOption("*", &marshalel{
			Marshaler:   &runtime.JSONPb{OrigName: true},
			contentType: "application/json; charset=UTF-8",
		}),
	)
	if err := protocol.RegisterDetectorHandlerServer(ctx, gw, s.Detector); err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/", gw)
	mux.Handle("/swagger/", http.StripPrefix("/swagger/",
		tools.SwaggerServer("/swagger/swagger.json", true),
	))
	mux.HandleFunc("/healthcheck", tools.HealthCheck)
	mux.Handle("/metrics", promhttp.Handler())

	h := middleware.HTTPContextWrapper(mux, s.ContextWrap)

	srv := &http.Server{Addr: address, Handler: h}
	go func() {
		<-ctx.Done()
		s.Logger.Info("Shutting down the HTTP server")
		if err := srv.Shutdown(context.Background()); err != nil {
			s.Logger.Error("Failed to shutdown HTTP server", zap.Error(err))
		}
	}()

	s.Logger.Info(fmt.Sprintf("Starting listening at %s", address))
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		s.Logger.Error("Failed to listen and serve", zap.Error(err))
		return err
	}
	return nil
}

func (s *GRPCServer) contextWrapFnk() func(ctx context.Context) (context.Context, error) {
	if s.ContextWrap == nil {
		return func(ctx context.Context) (context.Context, error) {
			return ctx, nil
		}
	}
	return func(ctx context.Context) (context.Context, error) {
		return s.ContextWrap(ctx), nil
	}
}

func parseNetwork(uri string) (network, address string) {
	if !strings.Contains(uri, ":") {
		return "tcp", uri
	}
	addr := strings.SplitN(uri, ":", 2)
	return addr[0], strings.Trim(addr[1], "/")
}

func loadCreds(crt, key string) (credentials.TransportCredentials, error) {
	if crt == `` {
		return insecure.NewCredentials(), nil
	}
	return credentials.NewServerTLSFromFile(crt, key)
}

type marshalel struct {
	contentType string
	runtime.Marshaler
}

func (m *marshalel) ContentType() string { return m.contentType }
