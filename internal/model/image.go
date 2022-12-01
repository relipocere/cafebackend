package model

// Image info is the metadata of an image.
type ImageMeta struct {
	ID            string
	OwnerUsername string
	Size          int64
	ContentType   string
}
