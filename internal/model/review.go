package model

// Review review of a store.
type Review struct {
	ID int64
	ReviewToCreate
}

// Review struct for creation of a store review.
type ReviewToCreate struct {
	AuthorUsername string
	StoreID        int64
	Rating         int64
	Commentary     string
}

// ReviewFilter filter to search reviews with.
type ReviewFilter struct {
	StoreIDs        []int64
	AuthorUsernames []string
	Rating          *IntRange
}
