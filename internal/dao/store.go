package dao

import (
	"context"

	logrpc "github.com/izaakdale/service-logrpc/api/v1"
)

type LogStore interface {
	Insert(context.Context, string, string, string, logrpc.Level) (string, error)
	Fetch(context.Context) ([]*logrpc.LogRecord, error)
}
