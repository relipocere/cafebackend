package image

import "github.com/relipocere/cafebackend/internal/model"

type Repo struct{}

func NewRepo() *Repo {
	return &Repo{}
}

type imageDTO struct {
	ID            string `db:"id"`
	OwnerUsername string `db:"owner_username"`
	Size          int64  `db:"byte_size"`
	ContentType   string `db:"content_type"`
}

func mapToImageMeta(dto imageDTO) model.ImageMeta {
	return model.ImageMeta{
		ID:            dto.ID,
		OwnerUsername: dto.OwnerUsername,
		Size:          dto.Size,
		ContentType:   dto.ContentType,
	}
}
