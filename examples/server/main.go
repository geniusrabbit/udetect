package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/demdxx/goconfig"
	"github.com/sspserver/udetect/examples/server/api"
	"github.com/sspserver/udetect/examples/server/context/config"
	"github.com/sspserver/udetect/examples/server/context/ctxlogger"
	"github.com/sspserver/udetect/examples/server/grpcserver"
)

var (
	buildCommit = ""
	appVersion  = "develop"
	buildDate   = ""
	conf        config.ConfigType
)

func main() {
	var wg sync.WaitGroup
	fatalError(goconfig.Load(&conf), "load config:")

	if conf.IsDebug() {
		fmt.Println(conf)
	}

	// Init new logger object
	loggerObj, err := newLogger(conf.IsDebug(), conf.LogLevel, zap.Fields(
		zap.String("commit", buildCommit),
		zap.String("version", appVersion),
		zap.String("build_date", buildDate),
	))

	fatalError(err, "init logger")

	// Register global logger
	zap.ReplaceGlobals(loggerObj)

	// Define cancelation context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	ctx = ctxlogger.WithLogger(ctx, loggerObj)
	defer cancel()

	// Init server config
	srv := grpcserver.GRPCServer{
		RequestTimeout: conf.Server.GRPC.Timeout,
		Logger:         loggerObj,
		CertFile:       conf.Server.GRPC.CertFile,
		KeyFile:        conf.Server.GRPC.KeyFile,
		Detector:       &api.Detector{},
	}

	if conf.Server.GRPC.Listen != "" {
		wg.Add(1)
		go func() {
			defer func() { wg.Done() }()
			fatalError(srv.RunGRPC(ctx, conf.Server.GRPC.Listen), "GRPC server")
		}()
	}
	if conf.Server.HTTP.Listen != "" {
		wg.Add(1)
		go func() {
			defer func() { wg.Done() }()
			fatalError(srv.RunHTTP(ctx, conf.Server.HTTP.Listen), "HTTP server")
		}()
	}

	wg.Wait()
}

func newLogger(debug bool, loglevel string, options ...zap.Option) (logger *zap.Logger, err error) {
	if debug {
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		return config.Build(options...)
	}

	var (
		level         zapcore.Level
		loggerEncoder = zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
		})
	)
	if err := level.UnmarshalText([]byte(loglevel)); err != nil {
		logger.Error("parse log level error", zap.Error(err))
	}
	core := zapcore.NewCore(loggerEncoder, os.Stdout, level)
	logger = zap.New(core, options...)

	return logger, nil
}

func fatalError(err error, msgs ...interface{}) {
	if err != nil {
		log.Fatalln(append(msgs, err)...)
	}
}
