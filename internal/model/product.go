package model

import "time"

// Product is a product in as store.
type Product struct {
	ID int64
	ProductCreate
}

// ProductCreate strut for product creation.
type ProductCreate struct {
	Name        string
	StoreID     int64
	PriceCents  int64
	Ingredients []string
	Calories    int64
	ImageID     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ProductFilter struct{
	StoreIDs []int64
	PriceCents *IntRange
	Calories *IntRange
}
