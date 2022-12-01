package image

import (
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/model"
)

// Create saves metadata of an image.
func (*Repo) Get(ctx context.Context, q database.Queryable, imageID string) (*model.ImageMeta, error) {
	qb := database.PSQL.
		Select("id", "owner_username", "byte_size", "content_type").
		From(database.TableImage).
		Where(sq.Eq{"id": imageID})

	var dto imageDTO
	err := q.Get(ctx, &dto, qb)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("get image: %w", err)
	}

	image := mapToImageMeta(dto)
	return &image, nil
}
