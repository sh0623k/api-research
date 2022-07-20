package server

import (
	"context"
	"errors"
	"strconv"
	todo "web-service/generated/grpc/todo/v1"
	"web-service/pkg/entities"
	"web-service/pkg/interfaces/grpc/todo/v1/converter/entity"

	"google.golang.org/grpc"
)

type TodoServer struct {
	todo.UnimplementedTodoManagerServer
	todos *entities.Todos
}

func NewTodoServer() *TodoServer {
	return &TodoServer{
		UnimplementedTodoManagerServer: todo.UnimplementedTodoManagerServer{},
		todos:                          entities.NewTodos(),
	}
}

func (s *TodoServer) RegisterTodoManagerServer(server *grpc.Server) {
	todo.RegisterTodoManagerServer(server, s)
}

func (s *TodoServer) CreateTodo(
	_ context.Context, request *todo.CreateTodoRequest) (*todo.Todo, error) {
	id := strconv.Itoa(s.todos.NewID())
	todoEntity := &entities.Todo{
		ID:   id,
		Text: request.GetText(),
		Done: false,
		User: &entities.User{
			ID:   request.GetUserId(),
			Name: "User " + request.GetUserId(),
		},
	}
	s.todos.Todos[id] = todoEntity
	return entity.ConvertToProtobufTodo(todoEntity), nil
}

func (s *TodoServer) FetchTodos(
	request *todo.FetchTodosRequest, server todo.TodoManager_FetchTodosServer) error {
	intSlice := make([]int, request.FetchCount)
	for index := range intSlice {
		todoEntity, ok := s.todos.Todos[strconv.Itoa(index+1)]
		if !ok {
			return nil
		}
		err := server.Send(entity.ConvertToProtobufTodo(todoEntity))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *TodoServer) FetchTodo(
	_ context.Context, request *todo.FetchTodoRequest) (*todo.Todo, error) {
	todoEntity, ok := s.todos.Todos[request.GetId()]
	if !ok {
		return nil, errors.New("there is not the todo with the specified id")
	}
	return entity.ConvertToProtobufTodo(todoEntity), nil
}

func (s *TodoServer) DeleteTodo(
	_ context.Context, request *todo.DeleteTodoRequest) (*todo.Todo, error) {
	todoEntity, ok := s.todos.Todos[request.GetId()]
	if !ok {
		return nil, errors.New("there is not the todo with the specified id")
	}
	delete(s.todos.Todos, request.GetId())
	return entity.ConvertToProtobufTodo(todoEntity), nil
}
