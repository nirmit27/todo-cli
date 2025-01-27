package tasks

import "time"

const (
	TimeFormat = time.RFC1123
	StatusTodo = "todo"
	StatusInProgress = "in-progress"
	StatusDone = "done"
)

type Task struct {
	Id 			int			`json:"id"`
	Description string		`json:"description"`
	Status 		string		`json:"status"`
	CreatedAt 	time.Time	`json:"createdAt"`
	UpdatedAt 	time.Time	`json:"updatedAt"`
}
