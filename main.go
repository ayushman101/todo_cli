package main

import (
	"fmt"
	"github.com/ayushman101/todo_cli/todos"
	"os"
	"flag"
)

//TODO: more flags to Add:  delete, completed, list

func main(){
	
	var delInd , completeIndex *int
	var addVal string
	var list *bool
	var t_color *string

	flag.StringVar(&addVal,"add","", "Use this flag followed by name of the task under quotation marks")
	
	delInd= flag.Int("delete",0 ,"Give the index of the task to be deleted")
	
	completeIndex= flag.Int("completed",0 ,"Give the index of the task to be marked completed")

	list=flag.Bool("ls",false,"Use this flag for Looking at all the tasks")

	t_color =flag.String("tc","blue",`Use this flag to change table color. Mention name of the table which could be any of the following:
	-tc blue
	-tc red
	-tc gray
	-tc yellow
	-tc green
	-tc white`)
	flag.Parse();

	tl,err:=todos.ReadFromFile("file.json")

	if err!=nil{
		

		_,err=os.Create("file.json")

		if err!=nil {
			fmt.Println("Could not create file")
			os.Exit(1)
		}
	}



	switch{
		case *list:
			//todos.TableColor(todos.ColorRed)
			tl.Display()

		case addVal!="" :
			t,err:=todos.NewTask(addVal)
			if err!=nil {
				fmt.Println("1",err)
				os.Exit(1)
			}

			tl.AddTask(t)
			err=tl.SaveToFile("file.json")
			
			if err!=nil {
				fmt.Println("2",err)
				os.Exit(1)
			}

			fmt.Println(addVal," task added ")

		case *completeIndex!=0 :
			err= tl.ToggleComplete(*completeIndex)
			if err!=nil {
				fmt.Println(err)
				os.Exit(1)
			}

			err=tl.SaveToFile("file.json")
			if err!=nil {
				fmt.Println(err)
				os.Exit(1)
			}

			
			fmt.Println("Task ",*completeIndex, " marked complete")
		case *delInd!=0:
			err= tl.DeleteTask(*delInd)
			if err!=nil {
				fmt.Println(err)
				os.Exit(1)
			}

			err=tl.SaveToFile("file.json")
			if err!=nil {
				fmt.Println(err)
				os.Exit(1)
			}
			
			fmt.Println("Task ",*delInd, " deleted ")
		
		case *t_color!="":
			colormap:=map[string]string{
				"white":todos.ColorDefault,
				"red":todos.ColorRed,
				"blue":todos.ColorBlue,
				"gray":todos.ColorGray,
				"green":todos.ColorGreen,
				"yellow":todos.ColorYellow,
			}

			todos.TableColor(colormap[*t_color])
		default :
			
			fmt.Println("Default")
			//todos.TableColor(todos.ColorGreen)
			tl.Display()
	}
}
