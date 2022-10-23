package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/edgedb/edgedb-go"
	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/model"
)

// GetSession retrieves session by username.
func (r *Repo) GetSession(ctx context.Context, q database.Queryable, username string) (*model.Session, error) {
	query := `select User{
		session: {
			session_id,
			expires_at	
		}
	}
	filter .username=<str>$0`

	var dto userDto
	var edbErr edgedb.Error
	err := q.QuerySingle(ctx, query, &dto, username)
	if errors.As(err, &edbErr) && edbErr.Category(edgedb.NoDataError) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	if dto.Session.Missing() {
		return nil, nil
	}

	return &model.Session{
		SessionID: dto.Session.SessionID,
		Username:  username,
		ExpiresAt: dto.Session.ExpiresAt,
	}, nil
}
