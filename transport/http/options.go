package http

import (
	"net"
	"net/http"
	"time"
	"unsafe"
)

// Option of transport connection
type Option func(tr *Transport)

// WithGET method of request
func WithGET(tr *Transport) {
	tr.method = http.MethodGet
}

// WithPOST method of request
func WithPOST(tr *Transport) {
	tr.method = http.MethodPost
}

// WithRoundTripper custom interface
func WithRoundTripper(transport http.RoundTripper) Option {
	return func(tr *Transport) {
		tr.client.Transport = transport
	}
}

// WithTimeout sets request timeouts
func WithTimeout(timeout, keepAlive time.Duration) Option {
	return func(tr *Transport) {
		tr.client.Timeout = timeout

		if t := getTransport(tr.client.Transport); t != nil {
			t.DialContext = (&net.Dialer{
				Timeout:   timeout,
				KeepAlive: keepAlive,
				DualStack: true,
			}).DialContext
			tr.client.Transport = t
		}
	}
}

// WithMaxConn from the client
func WithMaxConn(connCount int) Option {
	return func(tr *Transport) {
		t := mustGetTransport(tr.client.Transport)
		t.MaxConnsPerHost = connCount
		t.MaxIdleConns = connCount
		t.MaxIdleConnsPerHost = connCount
		tr.client.Transport = t
	}
}

func mustGetTransport(transport http.RoundTripper) *http.Transport {
	t := getTransport(transport)
	if t == nil {
		panic("unsupported transport object")
	}
	return t
}

func getTransport(transport http.RoundTripper) *http.Transport {
	if t, _ := transport.(*http.Transport); t != nil {
		dt, _ := http.DefaultTransport.(*http.Transport)
		if dt != nil && (unsafe.Pointer)(t) != (unsafe.Pointer)(dt) {
			return t
		}
	}
	return &http.Transport{}
}
