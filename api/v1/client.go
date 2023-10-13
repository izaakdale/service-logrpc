package logrpc

import (
	context "context"
	"crypto/tls"
	"crypto/x509"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

// Write logs the message to the datastore. Context should have a value with the key TraceIDKey.
// See TraceIdMiddleware() to add a trace ID to each http request.
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
