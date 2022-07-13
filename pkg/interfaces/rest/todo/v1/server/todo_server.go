package server

import (
	"fmt"
	"net/http"
	"strconv"
	"web-service/generated/rest/types"
	"web-service/pkg/entities"

	"github.com/gin-gonic/gin"
)

type TodoServer struct {
	todos *entities.Todos
}

func NewTodoServer() *TodoServer {
	return &TodoServer{todos: entities.NewTodos()}
}

func (t *TodoServer) FetchTodos(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, t.todos)
}

func (t *TodoServer) FetchTodo(ctx *gin.Context, id int64) {
	todoWithID, ok := t.todos.Todos[strconv.Itoa(int(id))]
	if !ok {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, todoWithID)
}

func (t *TodoServer) CreateTodo(ctx *gin.Context) {
	var newTodo types.NewTodo

	if err := ctx.BindJSON(&newTodo); err != nil {
		return
	}
	todoEntity := &entities.Todo{
		ID:   strconv.Itoa(t.todos.NewID()),
		Text: newTodo.Text,
		Done: false,
		User: &entities.User{
			ID:   newTodo.UserID,
			Name: fmt.Sprintf("User %s", newTodo.UserID),
		},
	}
	t.todos.Todos[todoEntity.ID] = todoEntity
	ctx.IndentedJSON(http.StatusCreated, newTodo)
}

func (t *TodoServer) DeleteTodo(ctx *gin.Context, id int64) {
	todoWithID, ok := t.todos.Todos[strconv.Itoa(int(id))]
	if !ok {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	delete(t.todos.Todos, strconv.Itoa(int(id)))
	ctx.IndentedJSON(http.StatusOK, todoWithID)
}
