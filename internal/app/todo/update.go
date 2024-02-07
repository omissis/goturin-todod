package todo

import (
	"context"

	"github.com/google/uuid"
	api "github.com/omissis/goturin-todod/internal/api/model"
	db "github.com/omissis/goturin-todod/internal/db/model"
)

func NewUpdate(model *db.Queries) *Update {
	return &Update{
		model: model,
	}
}

type Update struct {
	model *db.Queries
}

func (l *Update) Execute(cw api.TodoUpdate, id string) (api.TodoRead, error) {
	if _, err := l.model.UpdateTodo(
		context.Background(),
		db.UpdateTodoParams{
			Uuid:        uuid.MustParse(id),
			Title:       cw.Title,
			Description: cw.Description,
		},
	); err != nil {
		return api.TodoRead{}, err
	}

	return api.TodoRead{
		Uuid:        id,
		Title:       cw.Title,
		Description: cw.Description,
	}, nil
}
