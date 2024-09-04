package models

type Todo struct {
	Type string
	PrivateTodo
	ProjectTodo
}

type PrivateTodo struct {
	Text   string
	Done   bool
	UserID uint
}

type ProjectTodo struct {
	Text      string
	Done      bool
	UserID    uint
	ProjectID uint
}
