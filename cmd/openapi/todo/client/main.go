package main

import (
	"api-research/generated/openapi/types"
	"api-research/pkg/gateways/gin"
	"context"
	"fmt"
	"strconv"
	"time"
)

const count = 1000

func main() {
	ctx := context.Background()
	gateway := gin.NewGateway()
	begin := time.Now()
	intSlice := make([]int, count)
	for index := range intSlice {
		_, err := gateway.CreateTodo(ctx, &types.NewTodo{Text: "test", UserID: strconv.Itoa(index)})
		if err != nil {
			fmt.Printf("%s", err.Error())
			return
		}
	}
	endCreateTime := time.Now().Sub(begin)
	fmt.Printf("created in %g seconds\n", endCreateTime.Seconds())

	begin = time.Now()
	for index := range intSlice {
		_, err := gateway.FetchTodo(ctx, strconv.Itoa(index+1))
		if err != nil {
			fmt.Printf("%s", err.Error())
			return
		}
	}
	endFindTime := time.Now().Sub(begin)
	fmt.Printf("fetched all todos in %g seconds\n", endFindTime.Seconds())

	begin = time.Now()
	todos, err := gateway.FetchTodos(ctx)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}
	endFindTodosTime := time.Now().Sub(begin)
	fmt.Printf("len: %d \n", len(todos.Todos))
	fmt.Printf("fetched todos slice in %g seconds\n", endFindTodosTime.Seconds())

	begin = time.Now()
	for index := range intSlice {
		_, err = gateway.DeleteTodo(ctx, strconv.Itoa(index+1))
		if err != nil {
			fmt.Printf("%s", err.Error())
			return
		}
	}
	endDeleteTime := time.Now().Sub(begin)
	fmt.Printf("deleted in %g seconds\n", endDeleteTime.Seconds())
}
