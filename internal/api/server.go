package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	api "github.com/omissis/goturin-todod/internal/api/model"
	"github.com/omissis/goturin-todod/internal/app/todo"
)

var (
	ErrListTodos  = errors.New("cannot list todos")
	ErrCreateTodo = errors.New("cannot create todo")
	ErrReadTodo   = errors.New("cannot find todo")
	ErrUpdateTodo = errors.New("cannot update todo")
	ErrDeleteTodo = errors.New("cannot delete todo")
)

func NewServer(
	echo *echo.Echo,
	address string,
	port uint16,
	listTodos *todo.List,
	createTodo *todo.Create,
	readTodo *todo.Read,
	updateTodo *todo.Update,
	deleteTodo *todo.Delete,
) *Server {
	return &Server{
		echo:       echo,
		address:    address,
		port:       port,
		listTodos:  listTodos,
		createTodo: createTodo,
		readTodo:   readTodo,
		updateTodo: updateTodo,
		deleteTodo: deleteTodo,
	}
}

type Server struct {
	api.ServerInterface

	echo    *echo.Echo
	address string
	port    uint16

	listTodos  *todo.List
	createTodo *todo.Create
	readTodo   *todo.Read
	updateTodo *todo.Update
	deleteTodo *todo.Delete
}

func (s *Server) Run() error {
	api.RegisterHandlers(s.echo, s)

	return s.echo.Start(fmt.Sprintf("%s:%d", s.address, s.port))
}

// TODO: add page filters
func (s *Server) ListTodos(ctx echo.Context) error {
	ws, err := s.listTodos.Execute()
	if err != nil {
		log.Println(err)

		return ErrListTodos
	}

	return ctx.JSON(http.StatusOK, ws)
}

func (s *Server) CreateTodo(ctx echo.Context) error {
	var w api.TodoCreate

	if err := ctx.Bind(&w); err != nil {
		log.Println(err)

		return ErrCreateTodo
	}

	// TODO: enable validation
	// if err := ctx.Validate(w); err != nil {
	// 	log.Println(err)

	// 	return ErrCreateTodo
	// }

	nw, err := s.createTodo.Execute(w)
	if err != nil {
		log.Println(err)

		return ErrCreateTodo
	}

	return ctx.JSON(http.StatusCreated, nw)
}

func (s *Server) DeleteTodo(ctx echo.Context, uuid string) error {
	if err := s.deleteTodo.Execute(uuid); err != nil {
		log.Println(err)

		return fmt.Errorf("%w '%s'", ErrDeleteTodo, uuid)
	}

	return ctx.JSON(http.StatusNoContent, nil)
}

func (s *Server) GetTodo(ctx echo.Context, uuid string) error {
	w, err := s.readTodo.Execute(uuid)
	if err != nil {
		log.Println(err)

		return fmt.Errorf("%w '%s'", ErrReadTodo, uuid)
	}

	return ctx.JSON(http.StatusOK, w)
}

func (s *Server) UpdateTodo(ctx echo.Context, id api.Uuid) error {
	w := api.TodoUpdate{}

	if err := ctx.Bind(&w); err != nil {
		log.Println(err)

		return fmt.Errorf("%w '%s'", ErrUpdateTodo, id)
	}

	// TODO: enable validation
	// if err := ctx.Validate(w); err != nil {
	// 	log.Println(err)

	// 	return fmt.Errorf("%s: %w", uuid, ErrUpdateTodo)
	// }

	nw, err := s.updateTodo.Execute(w, id)
	if err != nil {
		log.Println(err)

		return fmt.Errorf("%w '%s'", ErrUpdateTodo, id)
	}

	return ctx.JSON(http.StatusOK, nw)
}
