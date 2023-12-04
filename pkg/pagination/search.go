package pagination

import "fmt"

type SearchOptions struct {
	Field string
	Value string
}

func (s *SearchOptions) GetQuery() string {
	return fmt.Sprintf("lower(%s) LIKE lower('%%%s%%')", s.Field, s.Value)
}
