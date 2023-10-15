package main

import (
	"fmt"
//	"github.com/ayushman101/todo_cli/todos"
//	"os"
	"flag"
)

//TODO: more flags to Add:  delete, completed, list

func main(){
	
	var delInd , completeIndex *int
	var addVal string
	flag.StringVar(&addVal,"add","", "Use this flag followed by name of the task under quotation marks")
	
	delInd= flag.Int("delete",0 ,"Give the index of the task to be deleted")
	
	completeIndex= flag.Int("completed",0 ,"Give the index of the task to be marked completed")

	flag.Parse();

	fmt.Printf("%s\n",addVal)

	fmt.Printf("%d\n%d\n",*delInd,*completeIndex)
}
