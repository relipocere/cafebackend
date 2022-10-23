package user

import (
	"context"
	"fmt"

	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/model"
)

// CreateSession creates new session.
func (r *Repo) CreateSession(ctx context.Context, q database.Queryable, session model.Session) error {
	query := `update User
		filter .username = <str>$0
		set{
		session := (insert AuthSession{
			session_id := <str>$1,
			expires_at := <datetime>$2
		})
	}`

	err := q.Execute(ctx, query, session.Username, session.SessionID, session.ExpiresAt)
	if err != nil {
		return fmt.Errorf("execute: %w", err)
	}

	return nil
}
