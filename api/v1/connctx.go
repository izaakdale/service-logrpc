package logrpc

import (
	"context"
	"net"

	"github.com/google/uuid"
)

func ConnectionContext(ctx context.Context, c net.Conn) context.Context {
	return context.WithValue(ctx, TraceIDKey, uuid.NewString())
}
