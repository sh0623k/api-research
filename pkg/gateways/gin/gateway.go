package gin

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"web-service/errors"
	"web-service/pkg/entities"
)

type gateway struct {
	client  *http.Client
	baseURL string
}

func NewGateway() *gateway {
	return &gateway{
		client:  &http.Client{},
		baseURL: "http://localhost:8080",
	}
}

func (g *gateway) FetchTodos(ctx context.Context) (todos *entities.Todos, err error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/todos", g.baseURL), http.NoBody)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	var response *http.Response
	response, err = g.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = errors.CombineErrors(err, response.Body.Close())
	}()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	todos = &entities.Todos{}
	err = json.Unmarshal(body, todos)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("length: %d \n", len(todos.Todos))
	// fmt.Printf("%s \n", todos.Todos["1"].Text)

	return todos, nil
}

func (g *gateway) FetchTodo(ctx context.Context, id string) (todo *entities.Todo, err error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/todos/%s", g.baseURL, id), http.NoBody)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	var response *http.Response
	response, err = g.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = errors.CombineErrors(err, response.Body.Close())
	}()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	todo = &entities.Todo{}
	err = json.Unmarshal(body, todo)
	if err != nil {
		return nil, err
	}
	/*
		if todo != nil {
			fmt.Printf("%s \n", todo.Text)
		}
	*/

	return todo, nil
}

func (g *gateway) DeleteTodo(ctx context.Context, id string) (todo *entities.Todo, err error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodDelete, fmt.Sprintf("%s/todos/%s", g.baseURL, id), http.NoBody)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	var response *http.Response
	response, err = g.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = errors.CombineErrors(err, response.Body.Close())
	}()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	todo = &entities.Todo{}
	err = json.Unmarshal(body, todo)
	if err != nil {
		return nil, err
	}
	/*
		if todo != nil {
			fmt.Printf("%s \n", todo.Text)
		}
	*/

	return todo, nil
}

func (g *gateway) CreateTodo(ctx context.Context, newTodo *entities.TodoInput) (createdTodo *entities.Todo, err error) {
	todo := &entities.Todo{
		ID:   "5",
		Text: newTodo.Text(),
		User: &entities.User{
			ID:   newTodo.UserID(),
			Name: "user " + newTodo.UserID(),
		},
	}
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(todo)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("%s", buf.String())
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/todos", g.baseURL), &buf)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	var response *http.Response
	response, err = g.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = errors.CombineErrors(err, response.Body.Close())
	}()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	createdTodo = &entities.Todo{}
	err = json.Unmarshal(body, createdTodo)
	if err != nil {
		return nil, err
	}
	/*
		if createdTodo != nil {
			fmt.Printf("%s \n", createdTodo.Text)
		}
	*/

	return createdTodo, nil
}