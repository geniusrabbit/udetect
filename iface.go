package udetect

import (
	"context"
	"io"

	"github.com/geniusrabbit/udetect/protocol"
)

// Transport interface defenition
type Transport interface {
	io.Closer
	Detect(ctx context.Context, req *protocol.Request) (*protocol.Response, error)
}
