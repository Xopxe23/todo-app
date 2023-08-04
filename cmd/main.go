package main

import (
	"log"

	"github.com/Xopxe23/todo-app"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error ocured while running http server: %s", err.Error())
	}
}