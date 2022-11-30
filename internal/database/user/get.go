package user

import (
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/model"
)

// GetFilter defines by what field user will be filtered.
type GetFilter func(sq.SelectBuilder) sq.SelectBuilder

// GetByUsername filters user by username.
func GetByUsername(username string) GetFilter {
	return func(qb sq.SelectBuilder) sq.SelectBuilder {
		return qb.Where(sq.Eq{"username": username})
	}
}

// GetBySessionID filters user by sessionID.
func GetBySessionID(sessionID string) GetFilter {
	return func(qb sq.SelectBuilder) sq.SelectBuilder {
		return qb.Where(sq.Eq{"session_id": sessionID})
	}
}

// Get retrieves user by username.
func (r *Repo) Get(ctx context.Context, q database.Queryable, filter GetFilter) (*model.User, error) {
	qb := database.PSQL.
		Select(
			"username",
			"full_name",
			"kind",
			"password_hash",
			"salt",
			"created_at",
			"updated_at",
			"session_id",
			"session_expires_at",
		).
		From(database.TableUser)

	qb = filter(qb)

	var dto userDto
	err := q.Get(ctx, &dto, qb)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	user, err := mapToUser(dto)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
