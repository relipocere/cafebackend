package user

import (
	"context"
	"fmt"

	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/model"
)

// Repo repository for working with users.
type Repo struct {
}

// NewRepo creates Repo.
func NewRepo() *Repo {
	return &Repo{}
}

// Create inserts new user.
func (r *Repo) Create(ctx context.Context, q database.Queryable, user model.User) (string, error) {
	qb := database.PSQL.Insert(database.TableUser).
		Columns(
			"username",
			"full_name",
			"kind",
			"password_hash",
			"salt",
			"created_at",
			"updated_at",
		).
		Values(
			user.Username,
			user.FullName,
			user.Kind,
			user.PasswordHash,
			user.Salt,
			user.CreatedAt,
			user.UpdatedAt,
		)

	_, err := q.Exec(ctx, qb)
	if err != nil {
		return "", fmt.Errorf("insert: %w", err)
	}

	return user.Username, nil
}
