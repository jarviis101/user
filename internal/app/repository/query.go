package repository

import (
	"fmt"
	"strings"
)

type queryBuilder struct {
	conditions []string
	args       []interface{}
}

func (q *queryBuilder) addWhere(expr string, val interface{}) {
	q.conditions = append(q.conditions, fmt.Sprintf(expr, len(q.args)+1))
	q.args = append(q.args, val)
}

func (q *queryBuilder) clause() string {
	if len(q.conditions) == 0 {
		return ""
	}

	return " WHERE " + strings.Join(q.conditions, " AND ")
}

func (q *queryBuilder) pagination(limit, offset int) string {
	query := ""

	q.args = append(q.args, limit)
	query += fmt.Sprintf(" LIMIT $%d", len(q.args))

	q.args = append(q.args, offset)
	query += fmt.Sprintf(" OFFSET $%d", len(q.args))

	return query
}
