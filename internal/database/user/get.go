package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/edgedb/edgedb-go"
	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/model"
)

// Get retrieves user-handler by username.
func (r *Repo) Get(ctx context.Context, q database.Queryable, username string) (*model.User, error) {
	query := `select User{
		id,
		username,
		full_name,
		kind,
		password_hash,
		salt,
		created_at,
		updated_at,
	}
	filter .username=<str>$0
	`

	var dto userDto
	var edbErr edgedb.Error
	err := q.QuerySingle(ctx, query, &dto, username)
	if errors.As(err, &edbErr) && edbErr.Category(edgedb.NoDataError) {
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

// GetBySessionID gets user-handler by the sessionID.
func (r *Repo) GetBySessionID(ctx context.Context, q database.Queryable, sessionID string) (*model.User, error) {
	query := `select User{
		id,
		username,
		full_name,
		kind,
		password_hash,
		salt,
		created_at,
		updated_at,
	}
	filter .session.session_id=<str>$0
	`

	var dto userDto
	var edbErr edgedb.Error
	err := q.QuerySingle(ctx, query, &dto, sessionID)
	if errors.As(err, &edbErr) && edbErr.Category(edgedb.NoDataError) {
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
