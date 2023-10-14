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
		CompletedAt: time.Time{},
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
		fmt.Printf("%v\t\t%s\t\t%v\t\t%s\n",i+1,task.Name,task.Completed,task.CreatedAt.Format(time.RFC822))
	}

}


func (t *Todos) DeleteTask(index int) error {


	if index<1 || index>len(*t) {
		return errors.New("Invalid Index")
	}

	tl:=*t

	tl=append(tl[:index],tl[index+1:]...)

	*t=tl

	return nil
} 


func (t *Todos) ToggleComplete(index int) error {

	if index<1 || index>len(*t) {
		return errors.New("Invalid Index")
	}


	t[index-1].Completed=true

	t[index-1].CompletedAt=time.Now()

	return nil
}
