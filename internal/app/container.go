package app

import (
	"database/sql"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/omissis/goturin-todod/internal/api"
	"github.com/omissis/goturin-todod/internal/app/todo"
	db "github.com/omissis/goturin-todod/internal/db/model"
	"github.com/omissis/goturin-todod/internal/x/uuid"
)

const (
	apiServerPort = 8080
	databasePort  = 5432
)

var ErrCannotCreateContainer = fmt.Errorf("cannot create container")

type ContainerFactoryFunc func() (*Container, error)

func NewDefaultParameters() Parameters {
	return Parameters{
		APIServerHost:     "0.0.0.0",
		APIServerPort:     apiServerPort,
		APIAllowedOrigins: []string{"http://localhost:3000", "http://todod.dev"},
		DBName:            "todod",
		DBHost:            "0.0.0.0",
		DBPassword:        "todod",
		DBPort:            databasePort,
		DBSslMode:         "disable",
		DBUser:            "todod",
	}
}

type Parameters struct {
	APIServerHost     string
	APIServerPort     uint16
	APIAllowedOrigins []string

	DBName     string
	DBHost     string
	DBPassword string
	DBPort     uint16
	DBSslMode  string
	DBUser     string
}

type services struct {
	apiServer     *api.Server
	echo          *echo.Echo
	database      *sql.DB
	sqlcModel     *db.Queries
	uuidGenerator *uuid.GoogleGenerator
	listTodos     *todo.List
	createTodo    *todo.Create
	readTodo      *todo.Read
	updateTodo    *todo.Update
	deleteTodo    *todo.Delete
}

func NewContainer() *Container {
	return &Container{
		Parameters: NewDefaultParameters(),
	}
}

type Container struct {
	Parameters
	services
}

func (c *Container) APIServer() *api.Server {
	if c.apiServer == nil {
		c.apiServer = api.NewServer(
			c.Echo(),
			c.Parameters.APIServerHost,
			c.Parameters.APIServerPort,
			c.ListTodos(),
			c.CreateTodo(),
			c.ReadTodo(),
			c.UpdateTodo(),
			c.DeleteTodo(),
		)
	}

	return c.apiServer
}

func (c *Container) Echo() *echo.Echo {
	if c.echo == nil {
		c.echo = echo.New()
		if len(c.Parameters.APIAllowedOrigins) > 0 {
			c.echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
				AllowOrigins: c.Parameters.APIAllowedOrigins,
				AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
			}))
		}
	}

	return c.echo
}

func (c *Container) Database() *sql.DB {
	if c.database == nil {
		dsn := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			c.Parameters.DBHost,
			c.Parameters.DBPort,
			c.Parameters.DBUser,
			c.Parameters.DBPassword,
			c.Parameters.DBName,
			c.Parameters.DBSslMode,
		)

		database, err := sql.Open("postgres", dsn)
		if err != nil {
			panic(err)
		}

		c.database = database
	}

	return c.database
}

func (c *Container) UUIDGenerator() *uuid.GoogleGenerator {
	if c.uuidGenerator == nil {
		c.uuidGenerator = uuid.NewGoogleGenerator()
	}

	return c.uuidGenerator
}

func (c *Container) ListTodos() *todo.List {
	if c.listTodos == nil {
		c.listTodos = todo.NewList(c.SqlcModel())
	}

	return c.listTodos
}

func (c *Container) CreateTodo() *todo.Create {
	if c.createTodo == nil {
		c.createTodo = todo.NewCreate(c.SqlcModel(), c.UUIDGenerator())
	}

	return c.createTodo
}

func (c *Container) ReadTodo() *todo.Read {
	if c.readTodo == nil {
		c.readTodo = todo.NewRead(c.SqlcModel())
	}

	return c.readTodo
}

func (c *Container) UpdateTodo() *todo.Update {
	if c.updateTodo == nil {
		c.updateTodo = todo.NewUpdate(c.SqlcModel())
	}

	return c.updateTodo
}

func (c *Container) DeleteTodo() *todo.Delete {
	if c.deleteTodo == nil {
		c.deleteTodo = todo.NewDelete(c.SqlcModel())
	}

	return c.deleteTodo
}

func (c *Container) SqlcModel() *db.Queries {
	if c.sqlcModel == nil {
		c.sqlcModel = db.New(c.Database())
	}

	return c.sqlcModel
}
