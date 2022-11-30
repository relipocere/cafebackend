package user

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/model"
)

// SetSession sets user's session.
func (r *Repo) SetSession(ctx context.Context, q database.Queryable, username string, session model.Session) error {
	qb := database.PSQL.
		Update(database.TableUser).
		SetMap(map[string]interface{}{
			"session_id":         session.ID,
			"session_expires_at": session.ExpiresAt,
		}).
		Where(sq.Eq{"username": username})

	_, err := q.Exec(ctx, qb)
	if err != nil {
		return fmt.Errorf("set session: %w", err)
	}

	return nil
}
