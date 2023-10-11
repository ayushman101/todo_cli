package todos

import (
	"time"
	"errors"
	"fmt"
)


type Task struct{
	Name string
	Completed bool
	CreatedAt time.Time
	CompletedAt time.Time
}

func NewTask (name string) (Task,error) {

	if name==""{
		return Task{},errors.New("No name given")
	}

	return Task{
		Name: name,
		Completed: false,
		CreatedAt: time.Now(),
		CompletedAt: time.Now(),
	},nil
}

type Todos []Task


func (t *Todos) AddTask(task Task) {

	tl:=*t
	tl=append(tl,task)
	*t=tl
}


func (t *Todos) Display() {

	fmt.Println("S.No.\t\tName\t\tCompleted\t\tCreatedAt")
	for i,task:= range *t {
		fmt.Printf("%v\t\t%s\t\t%v\t\t%s",i+1,task.Name,task.Completed,task.CreatedAt.Format(time.RFC822))
	}

}

