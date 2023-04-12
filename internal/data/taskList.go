package data

import (
	"golang.org/seun/Todo/internal/models"
)

var TodoList = []models.Task{
	{
		Id:        "1",
		Name:      "Clean tables and Chairs",
		Completed: false,
	},
	{
		Id:        "2",
		Name:      "Sweep the Balcony",
		Completed: false,
	},
	{
		Id:        "3",
		Name:      "Cook Spaghtti",
		Completed: false,
	},
}
