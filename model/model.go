package model

type TodoList struct {
	ID    int
	Title string
	Todos []Todo
}

type Todo struct {
	ID   int
	Name string
	Done bool
}
