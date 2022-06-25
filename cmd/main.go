package main

import (
	"log"
	todo "myapp/pkg"
	handler "myapp/pkg/handlers"
)

func main() {
	server := new(todo.Server)
	handlers := new(handler.Handler)
	if err := server.Run("8001", handlers.InitRoutes()); err != nil {
		log.Fatal("Error")
	}
}
