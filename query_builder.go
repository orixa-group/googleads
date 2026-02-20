package googleads

import (
	"fmt"
	"strings"
)

type QueryBuilder[F ~func() string] struct {
	fields       []string
	resource     string
	whereClauses []F
}

func NewQueryBuilder[F ~func() string]() *QueryBuilder[F] {
	return &QueryBuilder[F]{}
}

func (b *QueryBuilder[F]) Select(fields ...string) *QueryBuilder[F] {
	b.fields = fields
	return b
}

func (b *QueryBuilder[F]) From(resource string) *QueryBuilder[F] {
	b.resource = resource
	return b
}

func (b *QueryBuilder[F]) Where(clauses ...F) *QueryBuilder[F] {
	b.whereClauses = clauses
	return b
}

func (b *QueryBuilder[F]) Build() string {
	query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(b.fields, ", "), b.resource)
	if len(b.whereClauses) > 0 {
		query += " WHERE " + strings.Join(Map(b.whereClauses, func(clause F) string {
			return clause()
		}), " AND ")
	}

	return query
}
