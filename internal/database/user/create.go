package user

import (
	"context"
	"fmt"

	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/model"
)

// Repo is the repository for working with users.
type Repo struct {
}

// NewRepo creates Repo.
func NewRepo() *Repo {
	return &Repo{}
}

// Create creates a new user-handler.
func (r *Repo) Create(ctx context.Context, q database.Queryable, user model.UserCreate) (string, error) {
	query := `insert User{
		username := <str>$0, 
		full_name := <str>$1,
		kind := <UserKind>$2,
		password_hash := <str>$3,
		salt := <str>$4,
		created_at := <datetime>$5,
		updated_at := <datetime>$6
	}`

	var dto userDto
	err := q.QuerySingle(
		ctx, query, &dto,
		user.Username,
		user.FullName,
		string(user.Kind),
		user.PasswordHash,
		user.Salt,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		return "", fmt.Errorf("insert: %w", err)
	}

	return dto.ID.String(), err
}
