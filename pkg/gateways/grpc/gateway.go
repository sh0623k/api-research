package grpc

import (
	todo "api-research/generated/grpc/todo/v1"
	"api-research/pkg/entities"
	converter "api-research/pkg/interfaces/grpc/todo/v1/converter/protobuf"
	"context"
	"io"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Gateway struct {
	client todo.TodoManagerClient
}

func NewGateway() (*Gateway, error) {
	clientConnection, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &Gateway{
		client: todo.NewTodoManagerClient(clientConnection),
	}, nil
}

func (g *Gateway) CreateTodo(ctx context.Context, todoInput *entities.TodoInput) (*entities.Todo, error) {
	request := &todo.CreateTodoRequest{Text: todoInput.Text(), UserId: todoInput.UserID()}
	response, err := g.client.CreateTodo(ctx, request)
	if err != nil {
		return nil, err
	}
	protobufUser := response.GetUser()
	return &entities.Todo{
		ID:   response.GetId().GetId(),
		Text: response.GetText(),
		Done: response.GetDone(),
		User: &entities.User{
			ID:   protobufUser.GetId(),
			Name: protobufUser.GetName(),
		},
	}, nil
}

func (g *Gateway) FetchTodos(ctx context.Context) ([]*entities.Todo, error) {
	todoSlice := make([]*entities.Todo, 0)
	request := &todo.FetchTodosRequest{FetchCount: int32(0)}
	responseStream, err := g.client.FetchTodos(ctx, request)
	if err != nil {
		return nil, err
	}
	for {
		todoProtobuf, receiveErr := responseStream.Recv()
		if receiveErr == io.EOF {
			break
		}
		if receiveErr != nil {
			return nil, receiveErr
		}
		todoSlice = append(todoSlice, converter.ConvertToTodoEntity(todoProtobuf))
	}
	return todoSlice, nil
}

func (g *Gateway) FetchTodo(ctx context.Context, id string) (*entities.Todo, error) {
	request := &todo.FetchTodoRequest{Id: id}
	response, err := g.client.FetchTodo(ctx, request)
	if err != nil {
		return nil, err
	}
	protobufUser := response.GetUser()
	return &entities.Todo{
		ID:   response.GetId().GetId(),
		Text: response.GetText(),
		Done: response.GetDone(),
		User: &entities.User{
			ID:   protobufUser.GetId(),
			Name: protobufUser.GetName(),
		},
	}, nil
}

func (g *Gateway) DeleteTodo(ctx context.Context, id string) (*entities.Todo, error) {
	request := &todo.DeleteTodoRequest{Id: id}
	response, err := g.client.DeleteTodo(ctx, request)
	if err != nil {
		return nil, err
	}
	protobufUser := response.GetUser()
	return &entities.Todo{
		ID:   response.GetId().GetId(),
		Text: response.GetText(),
		Done: response.GetDone(),
		User: &entities.User{
			ID:   protobufUser.GetId(),
			Name: protobufUser.GetName(),
		},
	}, nil
}
