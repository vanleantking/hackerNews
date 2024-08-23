package model

type Pagination struct {
	PageSize    int
	Offset      int
	OrderBy     string
	CurrentPage int
}
