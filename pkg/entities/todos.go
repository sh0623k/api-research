package entities

type Todos struct {
	Todos     map[string]*Todo `json:"todos"`
	currentID int
}

func NewTodos() *Todos {
	return &Todos{
		Todos:     make(map[string]*Todo, 0),
		currentID: 0,
	}
}

func (t *Todos) NewID() int {
	t.currentID++
	return t.currentID
}
