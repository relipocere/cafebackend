package database

import (
	"context"
	"errors"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/relipocere/cafebackend/internal/model"
	"go.uber.org/zap"
)

// PSQL query builder with postgres placeholder already set.
var PSQL = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

var likeQuerySanitizer = strings.NewReplacer(`*`, `\*`, `_`, `\_`)

func ApplyIntFilter(qb sq.SelectBuilder, field string, r model.IntRange) sq.SelectBuilder {
	if r.Start != nil {
		if r.StartExclusive {
			qb = qb.Where(sq.Gt{field: *r.Start})
		} else {
			qb = qb.Where(sq.GtOrEq{field: *r.Start})
		}
	}

	if r.End != nil {
		if r.EndExclusive {
			qb = qb.Where(sq.Lt{field: *r.End})
		} else {
			qb = qb.Where(sq.LtOrEq{field: *r.End})
		}
	}

	return qb
}

// PreventNullSlice converts nil pointer to empty slice.
func PreventNullSlice[T any](slice []T) []T {
	if slice == nil {
		return make([]T, 0)
	}

	return slice
}

func SanitizeLikeQuery(query string) string {
	return likeQuerySanitizer.Replace(query)
}

func RollbackTx(ctx context.Context, tx Tx, callerFuncName string) {
	err := tx.Rollback(ctx)
	if err != nil && !errors.Is(err, pgx.ErrTxClosed) {
		zap.S().Errorf("rollback tx in %s: %w", callerFuncName, err)
	}
}
