package googleads

import (
	"strings"
)

type Filter func() string

func buildQuery(query string, filters ...Filter) string {
	if len(filters) > 0 {
		query += " WHERE " + strings.Join(Map(filters, func(filter Filter) string {
			return filter()
		}), " AND ")
	}

	return query
}
