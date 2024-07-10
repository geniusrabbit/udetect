package grpc

import (
	"context"
	"io"
	"runtime"
	"sync"
	"time"

	"github.com/geniusrabbit/udetect/protocol"
	"github.com/hashicorp/go-multierror"
)

// APIClient alias
type DetectorClient = protocol.DetectorClient

// ConnectionPool provides function to work with client connections
type ConnectionPool interface {
	io.Closer
	Client() (*Transport, error)
	Detect(ctx context.Context, req *protocol.Request) (*protocol.Response, error)
}

type lazyTransportPool struct {
	iterator uint32
	mx       sync.Mutex

	ctx     context.Context
	address string
	options []Option

	connectionTimeout time.Duration

	pool []*Transport
}

// NewLazyTransportPool provides connection pool to the service
func NewLazyTransportPool(ctx context.Context, poolSize int, grpcAddress string,
	connTimeout time.Duration, options ...Option) ConnectionPool {
	if poolSize <= 0 {
		poolSize = runtime.NumCPU()
	}
	return &lazyTransportPool{
		ctx:               ctx,
		address:           grpcAddress,
		connectionTimeout: connTimeout,
		options:           options,
		pool:              make([]*Transport, poolSize),
	}
}

// Client returns the link for the next API connection
func (cl *lazyTransportPool) Client() (*Transport, error) {
	cl.mx.Lock()
	defer cl.mx.Unlock()
	idx := cl.iterator % uint32(len(cl.pool))
	cl.iterator++
	if cl.pool[idx] == nil {
		client, err := cl.newTr()
		if err != nil {
			return nil, err
		}
		cl.pool[idx] = client
	}
	return cl.pool[idx], nil
}

// Detect user information
func (cl *lazyTransportPool) Detect(ctx context.Context, req *protocol.Request) (*protocol.Response, error) {
	cli, err := cl.Client()
	if err != nil {
		return nil, err
	}
	return cli.Detect(ctx, req)
}

// Close all client connections
func (cl *lazyTransportPool) Close() (allErr error) {
	cl.mx.Lock()
	defer cl.mx.Unlock()
	for _, client := range cl.pool {
		if err := client.Close(); err != nil {
			allErr = multierror.Append(allErr, err)
		}
	}
	return allErr
}

func (cl *lazyTransportPool) newTr() (*Transport, error) {
	ctx := cl.ctx
	if cl.connectionTimeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, cl.connectionTimeout)
		defer cancel()
	}
	return NewTransport(ctx, cl.address, cl.options...)
}
