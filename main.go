package main

import (
	"fmt"
	"github.com/ayushman101/todo_cli/todos"
	"os"
)

func main(){

	t,err:= todos.NewTask("First Sample Task")

	if err!=nil {
		fmt.Printf("Error: %w",err)
		os.Exit(1)
	}
		
	fmt.Println(t)

}
