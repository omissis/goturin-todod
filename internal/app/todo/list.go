package todo

import (
	"context"

	api "github.com/omissis/goturin-todod/internal/api/model"
	db "github.com/omissis/goturin-todod/internal/db/model"
)

func NewList(model *db.Queries) *List {
	return &List{
		model: model,
	}
}

type List struct {
	model *db.Queries
}

func (l *List) Execute() ([]api.TodoRead, error) {
	ws, err := l.model.ListTodos(context.Background())
	if err != nil {
		return nil, err
	}

	todos := make([]api.TodoRead, len(ws))

	for i, w := range ws {
		todos[i], err = mapDbToApi(w)
		if err != nil {
			return nil, err
		}
	}

	return todos, nil
}
