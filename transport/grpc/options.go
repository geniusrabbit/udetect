package grpc

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

// Options of the client connect
type Options struct {
	// Secure connections
	Secure bool

	// GRPCAddress contains HOST:PORT or IP:PORT to connecto to Pagecache GRPC server
	GRPCAddress string

	// Connect Certificates .pem, .crt
	CrtFile     string
	KeyFile     string
	RootCrtFile string
	ServerName  string
}

// TransportCredentials init GRPC transport credentionals
func (opt *Options) TransportCredentials() (credentials.TransportCredentials, error) {
	if opt.CrtFile == `` {
		return insecure.NewCredentials(), nil
	}
	addr, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	// Load the client certificates from disk
	certificate, err := tls.LoadX509KeyPair(opt.CrtFile, opt.KeyFile)
	if err != nil {
		return nil, fmt.Errorf("could not load client key pair: %s", err)
	}

	// Create a certificate pool from the certificate authority
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(opt.RootCrtFile)
	if err != nil {
		return nil, fmt.Errorf("could not read ca certificate: %s", err)
	}

	// Append the certificates from the CA
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		return nil, errors.New("failed to append ca certs")
	}

	return credentials.NewTLS(&tls.Config{
		ServerName:         addr, // this is required!
		Certificates:       []tls.Certificate{certificate},
		RootCAs:            certPool,
		InsecureSkipVerify: true,
	}), nil
}

// Option configuration type
type Option func(*Options)

// WithSecure connection
func WithSecure(secure bool) Option {
	return func(opt *Options) {
		opt.Secure = secure
	}
}

// WithTransportCredentials inits secure certificates
func WithTransportCredentials(crt, key, ca, serverName string) Option {
	return func(opt *Options) {
		opt.CrtFile = crt
		opt.KeyFile = key
		opt.RootCrtFile = ca
		opt.ServerName = serverName
	}
}
