package models

type Task struct {
	Id			string	`json: "id"`
	Name		string	`json: "title"`
	Completed	bool	`json: "completed"`
}