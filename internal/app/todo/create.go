package todo

import (
	"context"

	"github.com/google/uuid"
	api "github.com/omissis/goturin-todod/internal/api/model"
	db "github.com/omissis/goturin-todod/internal/db/model"
	xuuid "github.com/omissis/goturin-todod/internal/x/uuid"
)

func NewCreate(model *db.Queries, uuidGenerator xuuid.Generator) *Create {
	return &Create{
		model:         model,
		uuidGenerator: uuidGenerator,
	}
}

type Create struct {
	model         *db.Queries
	uuidGenerator xuuid.Generator
}

func (l *Create) Execute(cw api.TodoCreate) (api.TodoRead, error) {
	id := l.uuidGenerator.Generate()

	if _, err := l.model.CreateTodo(
		context.Background(),
		db.CreateTodoParams{
			Uuid:        uuid.MustParse(id),
			Title:       cw.Title,
			Description: *cw.Description,
		},
	); err != nil {
		return api.TodoRead{}, err
	}

	return api.TodoRead{
		Uuid:        id,
		Title:       cw.Title,
		Description: *cw.Description,
	}, nil
}
