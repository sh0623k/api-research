package main

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"web-service/pkg/entities"
	"web-service/pkg/gateways/grpc"
)

const count = 1000

func main() {
	ctx := context.Background()
	gateway, gatewayErr := grpc.NewGateway()
	if gatewayErr != nil {
		return
	}
	begin := time.Now()
	intSlice := make([]int, count)
	for index := range intSlice {
		_, err := gateway.CreateTodo(ctx, entities.NewTodoInput("test", strconv.Itoa(index)))
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
	_, err := gateway.FetchTodos(ctx, count)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}
	// fmt.Printf("len: %d \n", len(todos))
	endFindTodosTime := time.Now().Sub(begin)
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
