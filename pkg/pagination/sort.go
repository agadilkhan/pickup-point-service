package pagination

import (
	"fmt"
)

type SortOptions struct {
	SortBy    string
	SortOrder string
}

func (so *SortOptions) GetOrderBy() string {
	return fmt.Sprintf("%s %s", so.SortBy, so.SortOrder)
}
