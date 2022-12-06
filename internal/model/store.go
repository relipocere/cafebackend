package model

import "time"

// Affordability rating of the store
type Affordability string

// Affordability enum
const (
	AffordabilityCheap      Affordability = "Cheap"
	AffordabilityAffordable Affordability = "Affordable"
	AffordabilityExpensive  Affordability = "Expensive"
)

var validAffordabilityKinds = map[Affordability]struct{}{
	AffordabilityCheap:      {},
	AffordabilityAffordable: {},
	AffordabilityExpensive:  {},
}

// IsValid checks whether the current value is valid.
func (a Affordability) IsValid() bool {
	_, ok := validAffordabilityKinds[a]
	return ok
}

// Cuisine type of the store
type Cuisine string

// Cuisine enum
const (
	CuisineAmerican Cuisine = "American"
	CuisineAsian    Cuisine = "Asian"
	CuisineEuropean Cuisine = "European"
)

var validCuisineKinds = map[Cuisine]struct{}{
	CuisineAmerican: {},
	CuisineAsian:    {},
	CuisineEuropean: {},
}

// IsValid checks whether the current value is valid.
func (c Cuisine) IsValid() bool {
	_, ok := validCuisineKinds[c]
	return ok
}

// Store is the restaurant/cafe.
type Store struct {
	ID int64
	StoreCreate
}

// Store the struct for the store creation.
type StoreCreate struct {
	Title           string
	Affordability   Affordability
	Cuisine         Cuisine
	OwnerUsername   string
	ImageID         string
	AverageRating   float64
	NumberOfReviews int64
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type StoreFilter struct {
	TitleQuery     *string
	AverageRating  *IntRange
	OwnerUsernames []string
	Affordability  []Affordability
	Cuisines       []Cuisine
}
