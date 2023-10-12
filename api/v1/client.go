package logrpc

import (
	context "context"
	"crypto/tls"
	"crypto/x509"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type ctxKey int32

const (
	TraceIDKey ctxKey = iota
)

type Client struct {
	name string
	LoggingServiceClient
}

func NewClient(name, dialAddr string, crt tls.Certificate, rootCAs *x509.CertPool) (*Client, error) {
	cfg := tls.Config{
		Certificates: []tls.Certificate{crt},
		RootCAs:      rootCAs,
	}
	conn, err := grpc.Dial(dialAddr, grpc.WithTransportCredentials(credentials.NewTLS(&cfg)))
	if err != nil {
		return nil, err
	}
	return &Client{name, NewLoggingServiceClient(conn)}, nil
}

func (c *Client) Write(ctx context.Context, message string, logLevel Level) error {
	trace, ok := ctx.Value(TraceIDKey).(string)
	if !ok {
		c.Log(ctx, &LogRecord{
			Service:  c.name,
			Message:  "trace id was missing from conext",
			LogLevel: Level_ERROR,
		})
	}

	_, err := c.Log(ctx, &LogRecord{
		Service:  c.name,
		Message:  message,
		LogLevel: logLevel,
		TraceId:  trace,
	})

	return err
}
