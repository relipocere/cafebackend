package image

import (
	"context"

	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/model"
)

// Create saves metadata of an image.
func (*Repo) Create(ctx context.Context, q database.Queryable, image model.ImageMeta) error {
	qb := database.PSQL.
		Insert(database.TableImage).
		Columns("id", "owner_username", "byte_size", "content_type").
		Values(image.ID, image.OwnerUsername, image.Size, image.ContentType)

	_, err := q.Exec(ctx, qb)
	if err != nil {
		return err
	}

	return nil
}
