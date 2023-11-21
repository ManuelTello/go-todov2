package model

import "time"

type Task struct {
	Todo string `json:"todo"`
}

type Board struct {
	Name    string
	Created time.Time
	Tasks   []Task
}

type Workspace struct {
	Boards []Board
}

func CreateTask(todo string) Task {
	var task Task = Task{
		Todo: todo,
	}

	return task
}
