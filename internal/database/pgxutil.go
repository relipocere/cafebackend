package database

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// pgxUtil wrapper that implements PGX interface.
type pgxUtil struct {
	pool *pgxpool.Pool
}

// NewPGX creates new pgxUtil.
func NewPGX(pool *pgxpool.Pool) PGX {
	return &pgxUtil{pool: pool}
}

// BeginTx starts transaction.
func (p *pgxUtil) BeginTx(ctx context.Context, txOptions *pgx.TxOptions) (Tx, error) {
	var txOpts pgx.TxOptions
	if txOptions != nil {
		txOpts = *txOptions
	}

	tx, err := p.pool.BeginTx(ctx, txOpts)
	if err != nil {
		return nil, fmt.Errorf("не удалось начать транзакцию: %w", err)
	}

	return &txUtil{pgxTx: tx}, nil
}

// Exec executes query.
func (p *pgxUtil) Exec(ctx context.Context, sqlizer Sqlizer) (pgconn.CommandTag, error) {
	return execFn(ctx, p.pool, sqlizer)
}

// Select scans multiple rows. If there are no rows returns nil.
func (p *pgxUtil) Select(ctx context.Context, dst interface{}, sqlizer Sqlizer) error {
	return selectFn(ctx, p.pool, dst, sqlizer)
}

// Get scans single row. If there are no rows returns pgx.ErrNoRows error.
func (p *pgxUtil) Get(ctx context.Context, dst interface{}, sqlizer Sqlizer) error {
	return getFn(ctx, p.pool, dst, sqlizer)
}

// Tx wrapper for transaction that implements Queryable interface.
type txUtil struct {
	pgxTx pgx.Tx
}

// Exec executes query.
func (t *txUtil) Exec(ctx context.Context, sqlizer Sqlizer) (pgconn.CommandTag, error) {
	return execFn(ctx, t.pgxTx, sqlizer)
}

// Select scans multiple rows. If there are no rows returns nil.
func (t *txUtil) Select(ctx context.Context, dst interface{}, sqlizer Sqlizer) error {
	return selectFn(ctx, t.pgxTx, dst, sqlizer)
}

// Get scans single row. If there are no rows returns pgx.ErrNoRows error.
func (t *txUtil) Get(ctx context.Context, dst interface{}, sqlizer Sqlizer) error {
	return getFn(ctx, t.pgxTx, dst, sqlizer)
}

// Commit commits transaction.
func (t *txUtil) Commit(ctx context.Context) error {
	return t.pgxTx.Commit(ctx)
}

// Rollback cancels transaction.
func (t *txUtil) Rollback(ctx context.Context) error {
	return t.pgxTx.Rollback(ctx)
}

func execFn(ctx context.Context, e execer, sqlizer Sqlizer) (pgconn.CommandTag, error) {
	query, args, err := sqlizer.ToSql()
	if err != nil {
		return pgconn.CommandTag{}, fmt.Errorf("ToSql: %w", err)
	}

	return e.Exec(ctx, query, args...)
}

func selectFn(ctx context.Context, q pgxscan.Querier, dst interface{}, sqlizer Sqlizer) error {
	query, args, err := sqlizer.ToSql()
	if err != nil {
		return fmt.Errorf("ToSql: %w", err)
	}

	return pgxscan.Select(ctx, q, dst, query, args...)
}

func getFn(ctx context.Context, q pgxscan.Querier, dst interface{}, sqlizer Sqlizer) error {
	query, args, err := sqlizer.ToSql()
	if err != nil {
		return fmt.Errorf("ToSql: %w", err)
	}

	return pgxscan.Get(ctx, q, dst, query, args...)
}
