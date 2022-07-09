package resolvers

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos  map[string]interface{}
	todoID *todoID
}

func NewResolver() *Resolver {
	return &Resolver{
		todos:  make(map[string]interface{}),
		todoID: newTodoID(),
	}
}
