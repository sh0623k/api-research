package main

import (
	"net/http"
	"strconv"
	"web-service/pkg/entities"

	"github.com/gin-gonic/gin"
)

var (
	todos = entities.NewTodos()
	/*
		todos = entities.Todos{
			Todos: map[string]*entities.Todo{
				"1": {ID: "1", Text: "Todo 1", Done: false, User: &entities.User{ID: "1", Name: "Jotaro"}},
				"2": {ID: "2", Text: "Todo 2", Done: false, User: &entities.User{ID: "2", Name: "Joseph"}},
				"3": {ID: "3", Text: "Todo 3", Done: false, User: &entities.User{ID: "3", Name: "Avdol"}},
			},
		}
	*/
)

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodoByID)
	router.POST("/todos", postTodos)
	router.DELETE("/todos/:id", deleteTodos)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func getTodoByID(c *gin.Context) {
	id := c.Param("id")
	todoWithID, ok := todos.Todos[id]
	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, todoWithID)
}

func postTodos(c *gin.Context) {
	var newTodo entities.Todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}
	newTodo.ID = strconv.Itoa(todos.NewID())
	todos.Todos[newTodo.ID] = &newTodo
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func deleteTodos(c *gin.Context) {
	id := c.Param("id")
	todoWithID, ok := todos.Todos[id]
	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	delete(todos.Todos, id)
	c.IndentedJSON(http.StatusOK, todoWithID)
}
