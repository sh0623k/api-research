package entities

type TodoInput struct {
	text   string
	userID string
}

func NewTodoInput(text, userID string) *TodoInput {
	return &TodoInput{
		text:   text,
		userID: userID,
	}
}

func (t *TodoInput) Text() string {
	return t.text
}

func (t *TodoInput) UserID() string {
	return t.userID
}
