package todo

import (
	"context"

	"github.com/google/uuid"
	api "github.com/omissis/goturin-todod/internal/api/model"
	db "github.com/omissis/goturin-todod/internal/db/model"
)

func NewRead(model *db.Queries) *Read {
	return &Read{
		model: model,
	}
}

type Read struct {
	model *db.Queries
}

func (r *Read) Execute(id string) (api.TodoRead, error) {
	ws, err := r.model.GetTodo(context.Background(), uuid.MustParse(id))
	if err != nil {
		return api.TodoRead{}, err
	}

	todo, err := mapDbToApi(ws)
	if err != nil {
		return api.TodoRead{}, err
	}

	return todo, nil
}
