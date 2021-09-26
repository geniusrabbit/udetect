package grpc

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/sspserver/udetect/protocol"
)

type Transport struct {
	protocol.DetectorClient
	cliConn *grpc.ClientConn

	options Options
}

// NewTransport connector pool
func NewTransport(ctx context.Context, grpcAddress string, options ...Option) (*Transport, error) {
	opts := Options{
		GRPCAddress: grpcAddress,
	}
	for _, opt := range options {
		opt(&opts)
	}

	grpcClient := &Transport{
		options: opts,
	}

	// Init certification
	creds, err := opts.TransportCredentials()
	if err != nil {
		return nil, errors.Wrap(err, `GRPC transport credentials`)
	}

	// Establich new connection and create GRPC API client
	grpcClient.cliConn, grpcClient.DetectorClient, err = newGRPCClient(ctx, opts.GRPCAddress, []grpc.DialOption{
		// DialOption which configures a connection level security credentials
		grpc.WithTransportCredentials(creds),
		// Dial blocks until the underlying connection is up
		grpc.WithBlock(),
	}...)

	if err != nil {
		return nil, errors.Wrap(err, `create new API client`)
	}
	return grpcClient, nil
}

// Detect user information
func (tr *Transport) Detect(ctx context.Context, req *protocol.Request) (*protocol.Response, error) {
	return tr.DetectorClient.Detect(ctx, req)
}

// Close GRPC connection to API
func (tr *Transport) Close() error {
	if tr.cliConn != nil {
		return tr.cliConn.Close()
	}
	return nil
}

func newGRPCClient(ctx context.Context, address string, opts ...grpc.DialOption) (*grpc.ClientConn, protocol.DetectorClient, error) {
	if !strings.Contains(address, "://") {
		address = "tcp://" + address
	}
	addUrl, err := url.Parse(address)
	if err != nil {
		return nil, nil, err
	}
	conn, err := dial(ctx, addUrl.Scheme, addUrl.Host, opts...)
	if err != nil {
		return nil, nil, err
	}
	return conn, protocol.NewDetectorClient(conn), nil
}

///////////////////////////////////////////////////////////////////////////////
/// Dialers
///////////////////////////////////////////////////////////////////////////////

func dial(ctx context.Context, network, addr string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	switch network {
	case "tcp", "grpc":
		return dialTCP(ctx, addr, opts...)
	case "dns":
		return grpc.DialContext(ctx, addr, opts...)
	case "unix":
		return dialUnix(ctx, addr, opts...)
	default:
		return nil, fmt.Errorf("unsupported network type %q", network)
	}
}

// dialTCP creates a client connection via TCP.
// "addr" must be a valid TCP address with a port number.
func dialTCP(ctx context.Context, addr string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}
	if net.ParseIP(host) == nil {
		ip, err := net.ResolveIPAddr("ip", host)
		if err != nil {
			return nil, err
		}
		addr = ip.String() + ":" + port
	}
	return grpc.DialContext(ctx, addr, opts...)
}

// dialUnix creates a client connection via a unix domain socket.
// "addr" must be a valid path to the socket.
func dialUnix(ctx context.Context, addr string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	return grpc.DialContext(ctx, addr, append(opts,
		grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {
			if deadline, ok := ctx.Deadline(); ok {
				return net.DialTimeout("unix", addr, time.Until(deadline))
			}
			return net.DialTimeout("unix", addr, 0)
		}))...)
}
