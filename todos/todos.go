package todos

import (
	"time"
	"errors"
	"fmt"
	"encoding/json"
	"os"
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

	fmt.Println("S.No.\t\tName\t\t\tCompleted\t\tCreatedAt\t\tCompletedAt")
	for i,task:= range *t {
		fmt.Printf("%v\t\t%s\t\t%v\t\t%s\t\t%s\n",i+1,task.Name,task.Completed,task.CreatedAt.Format(time.RFC822), task.CompletedAt.Format(time.RFC822))
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

	tl:=*t

	tl[index-1].Completed=true

	tl[index-1].CompletedAt=time.Now()

	*t=tl

	return nil
}

func (t *Todos) SaveToFile(filename string) error {

	tl:=*t

	jsonData,err:= json.Marshal(tl)

	if err!=nil {
		
		return errors.New("Failed to Parse slice to Json")

	}
	
	//creating the file for writing json
	file,err:= os.Create(filename) 

	if err!=nil {
		
		return errors.New("Failed to open the file")
	}

	_,err= file.WriteString(string(jsonData))

	if err!=nil {
		return errors.New("Failed to write to json file")
	}
	
	return nil
}

func ReadFromFile(filename string) (*Todos, error) {

	data,err:= os.ReadFile(filename)

	if err!=nil {
		return nil,errors.New("File failed to open")
	}

	var t *Todos

	err= json.Unmarshal(data,&t)
	
	if err!=nil {
		return nil,fmt.Errorf("Failed to Unmarshal: %w",err)
	}

	return t,nil;
}


