package main

import (
	"api-research/generated/graphql/client"
	"api-research/generated/openapi/types"
	"api-research/pkg/entities"
	"api-research/pkg/gateways/gin"
	"api-research/pkg/gateways/grpc"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Khan/genqlient/graphql"
)

const count = 1000

func main() {
	ctx := context.Background()
	intSlice := make([]int, count)

	graphqlClient := graphql.NewClient("http://localhost:8080/query", http.DefaultClient)
	for index := range intSlice {
		_, createErr := client.CreateTodo(ctx, graphqlClient, fmt.Sprintf("todo%d", index), strconv.Itoa(index))
		if createErr != nil {
			fmt.Printf("%s", createErr.Error())
			return
		}
	}
	graphqlTodos, err := client.FindTodos(ctx, graphqlClient)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}
	fmt.Printf("len: %d \n", len(graphqlTodos.Todos))

	grpcGateway, err := grpc.NewGateway()
	if err != nil {
		return
	}
	for index := range intSlice {
		_, createErr := grpcGateway.CreateTodo(ctx, entities.NewTodoInput("test", strconv.Itoa(index)))
		if createErr != nil {
			fmt.Printf("%s", createErr.Error())
			return
		}
	}
	grpcTodos, err := grpcGateway.FetchTodos(ctx)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}
	fmt.Printf("len: %d \n", len(grpcTodos))

	ginGateway := gin.NewGateway()
	for index := range intSlice {
		_, createErr := ginGateway.CreateTodo(ctx, &types.NewTodo{Text: "test", UserID: strconv.Itoa(index)})
		if createErr != nil {
			fmt.Printf("%s", createErr.Error())
			return
		}
	}
	oapiTodos, err := ginGateway.FetchTodos(ctx)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}
	fmt.Printf("len: %d \n", len(oapiTodos.Todos))
}
