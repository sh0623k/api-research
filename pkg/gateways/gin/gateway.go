package gin

import (
	"api-research/errors"
	"api-research/generated/openapi/types"
	"api-research/pkg/entities"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type gateway struct {
	client  *http.Client
	baseURL string
}

func NewGateway() *gateway {
	return &gateway{
		client:  &http.Client{},
		baseURL: "http://localhost:3000",
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

	return todo, nil
}

func (g *gateway) CreateTodo(ctx context.Context, newTodo *types.NewTodo) (createdTodo *entities.Todo, err error) {
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(newTodo)
	if err != nil {
		return nil, err
	}
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

	return createdTodo, nil
}
