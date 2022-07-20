package main

import (
	"api-research/generated/graphql/client"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Khan/genqlient/graphql"
)

const count = 1000

func main() {
	graphqlClient := graphql.NewClient("http://localhost:8080/query", http.DefaultClient)
	ctx := context.Background()
	begin := time.Now()
	intSlice := make([]int, count)
	for index := range intSlice {
		_, err := client.CreateTodo(ctx, graphqlClient, fmt.Sprintf("todo%d", index), strconv.Itoa(index))
		if err != nil {
			fmt.Printf("%s", err.Error())
			return
		}
	}
	endCreateTime := time.Now().Sub(begin)
	fmt.Printf("created in %g seconds\n", endCreateTime.Seconds())

	begin = time.Now()
	for index := range intSlice {
		_, findErr := client.FindTodo(ctx, graphqlClient, strconv.Itoa(index+1))
		if findErr != nil {
			fmt.Printf("%s", findErr.Error())
			return
		}
	}
	endFindTime := time.Now().Sub(begin)
	fmt.Printf("fetched all todos in %g seconds\n", endFindTime.Seconds())

	begin = time.Now()
	todos, err := client.FindTodos(ctx, graphqlClient)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}
	endFindTodosTime := time.Now().Sub(begin)
	fmt.Printf("len: %d \n", len(todos.Todos))
	fmt.Printf("fetched todos slice in %g seconds\n", endFindTodosTime.Seconds())

	begin = time.Now()
	for index := range intSlice {
		_, deleteErr := client.DeleteTodo(ctx, graphqlClient, strconv.Itoa(index+1))
		if deleteErr != nil {
			fmt.Printf("%s", deleteErr.Error())
			return
		}
	}
	endDeleteTime := time.Now().Sub(begin)
	fmt.Printf("deleted in %g seconds\n", endDeleteTime.Seconds())
}
