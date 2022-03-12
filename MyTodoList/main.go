/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/majorchork/week4/mytodolist/MyTodoList/cmd"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile | log.Lshortfile)
	cmd.Execute()
}
