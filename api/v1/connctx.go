package logrpc

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type ctxKey int32

const (
	TraceIDKey ctxKey = iota
)

func TraceIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(context.WithValue(r.Context(), TraceIDKey, uuid.NewString()))
		next.ServeHTTP(w, r)
	})
}
