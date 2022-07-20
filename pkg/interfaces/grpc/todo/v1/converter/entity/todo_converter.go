package entity

import (
	protobuf "api-research/generated/grpc/todo/v1"
	"api-research/pkg/entities"
)

func ConvertToProtobufTodo(todoEntity *entities.Todo) *protobuf.Todo {
	userEntity := todoEntity.User
	return &protobuf.Todo{
		Id:   &protobuf.FetchTodoRequest{Id: todoEntity.ID},
		Text: todoEntity.Text,
		Done: todoEntity.Done,
		User: &protobuf.User{
			Id:   userEntity.ID,
			Name: userEntity.Name,
		},
	}
}
