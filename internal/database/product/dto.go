package product

import "time"

type Repo struct {
}

func NewRepo() *Repo {
	return &Repo{}
}

type productDTO struct {
	ID          int64     `db:"id"`
	Name        string    `db:"name"`
	StoreID     int64     `db:"store_id"`
	Ingerdients []string  `db:"ingredients"`
	Calories    int64     `db:"calories"`
	ImageID     string    `db:"image_id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
