package repository

import (
	"fmt"
	"strings"
)

type whereBuilder struct {
	conditions []string
	args       []interface{}
}

func (w *whereBuilder) add(expr string, val interface{}) {
	w.conditions = append(w.conditions, fmt.Sprintf(expr, len(w.args)+1))
	w.args = append(w.args, val)
}

func (w *whereBuilder) clause() string {
	if len(w.conditions) == 0 {
		return ""
	}

	return " WHERE " + strings.Join(w.conditions, " AND ")
}
