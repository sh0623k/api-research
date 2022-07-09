package resolvers

import "strconv"

type todoID struct {
	current int
}

func newTodoID() *todoID {
	return &todoID{current: 0}
}

func (t *todoID) newID() string {
	t.current++
	return strconv.Itoa(t.current)
}
