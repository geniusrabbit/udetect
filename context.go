package udetect

import "context"

var ctxCliKey = struct{ s string }{s: "udetect.cli"}

// FromContext returns client from context
func FromContext(ctx context.Context) *Client {
	return ctx.Value(ctxCliKey).(*Client)
}

// WithContext returns new context with client
func WithContext(ctx context.Context, cli *Client) context.Context {
	return context.WithValue(ctx, ctxCliKey, cli)
}
