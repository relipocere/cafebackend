package model

type IntRange struct {
	Start          *int64
	End            *int64
	StartExclusive bool
	EndExclusive   bool
}

type Pagination struct {
	Page         int64
	ItemsPerPage int64
}

func (p Pagination) Limit() uint64 {
	return uint64(p.ItemsPerPage)
}

func (p Pagination) Offset() uint64 {
	return uint64((p.Page - 1) * p.ItemsPerPage)
}
