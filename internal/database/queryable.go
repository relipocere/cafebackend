package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// PGX interface containing operations necessary to interact with the database.
type PGX interface {
	Queryable
	BeginTx(ctx context.Context, txOptions *pgx.TxOptions) (Tx, error)
}

// Tx transaction interface.
type Tx interface {
	Queryable
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

// Queryable interface containing operations necessary to query the database.
type Queryable interface {
	Exec(ctx context.Context, sqlizer Sqlizer) (pgconn.CommandTag, error)
	Get(ctx context.Context, dst interface{}, sqlizer Sqlizer) error
	Select(ctx context.Context, dst interface{}, sqlizer Sqlizer) error
}

// Sqlizer copy of sql.Sqlizer interface.
type Sqlizer interface {
	ToSql() (sql string, args []interface{}, err error)
}

type execer interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
}
