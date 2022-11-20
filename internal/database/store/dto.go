package store

import (
	"time"

	"github.com/edgedb/edgedb-go"
)

type storeDto struct {
	ID            edgedb.UUID `edgedb:"id"`
	Title         string      `edgedb:"title"`
	Affordability string      `edgedb:"affordability"`
	OwnerUsername string      `edgedb:"owner_username"`
	ImageID       string      `edgedb:"image_id"`
	CreatedAt     time.Time   `edgedb:"created_at"`
	UpdatedAt     time.Time   `edgedb:"updated_At"`
}
