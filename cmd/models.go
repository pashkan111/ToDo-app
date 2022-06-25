package main

type User struct {
	Id       int    `json: "id"`
	Name     string `json: "name"`
	Username string `json: "username"`
	Password string `json: "password"`
}

// type ToDoList struct {
// 	Id         int
// 	ToDoItemId int
// 	UserId     int
// }

type ToDoList struct {
	Id          int    `json: "id"`
	Title       string `json: "title"`
	Description string `json: "description"`
}

type ToDoListItem struct {
	Id          int    `json: "id"`
	Title       string `json: "title"`
	Description string `json: "description"`
	Done        bool   `json: "done"`
}
