package database

import (
	"context"

	"github.com/edgedb/edgedb-go"
)

// Edge database.
type Edge interface {
	Tx(ctx context.Context, action edgedb.TxBlock) error
	Queryable
}

// Queryable interface for querying the database.
type Queryable interface {
	Execute(ctx context.Context, cmd string, args ...interface{}) error
	Query(ctx context.Context, cmd string, out interface{}, args ...interface{}) error
	QuerySingle(ctx context.Context, cmd string, out interface{}, args ...interface{}) error
	QueryJSON(ctx context.Context, cmd string, out *[]byte, args ...interface{}) error
	QuerySingleJSON(ctx context.Context, cmd string, out interface{}, args ...interface{}) error
}
