package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"api-research/generated/graphql/model"
	"api-research/generated/graphql/server"
	"context"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	id := r.todoID.newID()
	todo := &model.Todo{
		Text: input.Text,
		ID:   id,
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.todos[id] = todo
	return todo, nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (*model.Todo, error) {
	todo, ok := r.todos[id]
	if !ok {
		return nil, nil
	}
	delete(r.todos, id)
	return todo.(*model.Todo), nil
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	todo, ok := r.todos[id]
	if !ok {
		return nil, nil
	}
	return todo.(*model.Todo), nil
}

func (r *queryResolver) Todos(ctx context.Context) (map[string]interface{}, error) {
	return r.todos, nil
}

// Mutation returns server.MutationResolver implementation.
func (r *Resolver) Mutation() server.MutationResolver { return &mutationResolver{r} }

// Query returns server.QueryResolver implementation.
func (r *Resolver) Query() server.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
