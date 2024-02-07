package todo

import (
	api "github.com/omissis/goturin-todod/internal/api/model"
	db "github.com/omissis/goturin-todod/internal/db/model"
)

func mapDbToApi(dbw db.Todo) (api.TodoRead, error) {
	return api.TodoRead{
		Uuid:        dbw.Uuid.String(),
		Title:       dbw.Title,
		Description: dbw.Description,
	}, nil
}
