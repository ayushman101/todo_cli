package main

import (
	"fmt"
//	"github.com/ayushman101/todo_cli/todos"
//	"os"
	"flag"
)

func main(){
	
	var addVal string
	flag.StringVar(&addVal,"add","Sample Task", "Use this flag followed by name of the task under quotation marks")

	flag.Parse();

	fmt.Printf("%s\n",addVal)
}
