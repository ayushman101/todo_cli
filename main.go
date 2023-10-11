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
		
	
	tl:=todos.Todos{}

	tl.AddTask(t)
	t,err= todos.NewTask("Second Task added")

	tl.AddTask(t)


	tl.Display()

	err=tl.DeleteTask(1)

	if err!=nil {

		fmt.Printf("Error: %w",err)
		os.Exit(1)
	}

	tl.Display()

}
