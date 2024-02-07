package todo

import (
	"context"

	"github.com/google/uuid"
	db "github.com/omissis/goturin-todod/internal/db/model"
)

func NewDelete(model *db.Queries) *Delete {
	return &Delete{
		model: model,
	}
}

type Delete struct {
	model *db.Queries
}

func (d *Delete) Execute(id string) error {
	if err := d.model.DeleteTodo(
		context.Background(),
		uuid.MustParse(id),
	); err != nil {
		return err
	}

	return nil
}
