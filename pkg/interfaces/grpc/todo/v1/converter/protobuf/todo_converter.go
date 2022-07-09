package protobuf

import (
	protobuf "web-service/generated/grpc/todo/v1"
	"web-service/pkg/entities"
)

func ConvertToTodoEntity(protobufTodo *protobuf.Todo) *entities.Todo {
	protobufUser := protobufTodo.GetUser()
	return &entities.Todo{
		ID:   protobufTodo.GetId().GetId(),
		Text: protobufTodo.GetText(),
		Done: protobufTodo.GetDone(),
		User: &entities.User{
			ID:   protobufUser.GetId(),
			Name: protobufUser.GetName(),
		},
	}
}
