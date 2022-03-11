/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"
	"github.com/majorchork/week4/mytodolist/MyTodoList/cmd"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile | log.Lshortfile)
	cmd.Execute()
	fmt.Println(`todolist CLI application will do the following:
· List: list out all tasks that are not done
· Cleanup: delete tasks that have been done
· Add: add task to the list of tasks
· Done: would mark a task as done
· Undone: would mark a task as undone`)
}
