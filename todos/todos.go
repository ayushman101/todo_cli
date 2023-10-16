package todos

import (
	"time"
	"errors"
	"fmt"
	"encoding/json"
	"os"
	"github.com/alexeyco/simpletable"
)

//ANSI Escape Codes for changing the colors
const (
	ColorDefault = "\x1b[39m"

	ColorRed   = "\x1b[91m"
	ColorGreen = "\x1b[32m"
	ColorBlue  = "\x1b[94m"
	ColorGray  = "\x1b[90m"
	ColorYellow= "\x1b[33m"
)


var ColorTable string=ColorBlue

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

//	fmt.Println("S.No.\t\tName\t\t\tCompleted\t\tCreatedAt\t\tCompletedAt")
//	for i,task:= range *t {
//		fmt.Printf("%v\t\t%s\t\t%v\t\t%s\t\t%s\n",i+1,task.Name,task.Completed,task.CreatedAt.Format(time.RFC822), task.CompletedAt.Format(time.RFC822))
	//}

	table:=simpletable.New()

	table.Header = &simpletable.Header{

		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: yellow("#")},
			{Align: simpletable.AlignCenter, Text: yellow("Task")},
			{Align: simpletable.AlignCenter, Text: yellow("Completed")},
			{Align: simpletable.AlignCenter, Text: yellow("CreatedAt")},
			{Align: simpletable.AlignCenter, Text: yellow("CompletedAt")},
		},
	}

	for i, row := range *t {
		
		color:=ColorRed
		if row.Completed{
			color=ColorGreen
		}

		r:=[]*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d",i)},
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%s%v%s", color, row.Name, ColorTable)},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%s%v%s",color,row.Completed,ColorTable)},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%s%s%s",color,row.CreatedAt.Format(time.RFC822),ColorTable)},

			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%s%s%s",color,row.CompletedAt.Format(time.RFC822),ColorTable)},

		}

		table.Body.Cells= append(table.Body.Cells, r)
	}
	table.SetStyle(simpletable.StyleRounded)
	fmt.Printf("%s",ColorTable)
	fmt.Println(table.String())

	fmt.Printf("%s",ColorDefault)
}


func yellow(text string) string {
	return fmt.Sprintf("%s%s%s",ColorYellow,text,ColorTable)
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


