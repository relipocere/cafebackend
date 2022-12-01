package image

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"github.com/relipocere/cafebackend/internal/auth"
	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/model"
)

type imageRepo interface {
	Create(ctx context.Context, q database.Queryable, image model.ImageMeta) error
}

type App struct {
	filesDir  string
	db        database.PGX
	imageRepo imageRepo
}

func NewApp(filesDir string, db database.PGX, imageRepo imageRepo) *App {
	return &App{
		filesDir:  filesDir,
		db:        db,
		imageRepo: imageRepo,
	}
}

func (a *App) UploadImage(ctx context.Context, image graphql.Upload) (string, error) {
	user, ok := ctx.Value(auth.User).(model.User)
	if !ok {
		return "", fmt.Errorf("no user in the context")
	}

	if user.Kind != model.UserKindBusiness {
		return "", model.Error{
			Code:    model.ErrorCodeBadRequest,
			Message: "The type of user account you have can't upload images",
		}
	}

	imageID := uuid.NewString()
	fileBytes, err := io.ReadAll(image.File)
	if err != nil {
		return "", fmt.Errorf("read all: %w", err)
	}

	filePath := filepath.Join(a.filesDir, imageID)
	err = os.WriteFile(filePath, fileBytes, fs.ModePerm)
	if err != nil {
		return "", fmt.Errorf("write file to '%s': %w", filePath, err)
	}

	err = a.imageRepo.Create(ctx, a.db, model.ImageMeta{
		ID:            imageID,
		OwnerUsername: user.Username,
		Size:          image.Size,
		ContentType:   image.ContentType,
	})
	if err != nil {
		return "", fmt.Errorf("save image meta: %w", err)
	}

	return imageID, nil
}
