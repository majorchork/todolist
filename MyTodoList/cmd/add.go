/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	cmd2 "github.com/majorchork/week4/mytodolist/MyTodoList/cmd/model"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
//type ListItem = struct {
//	SN     int
//	Item   string
//	Status bool
//}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds task to the list",
	Long: `Use this command to add task you want to carry out:

		this function add's items to your todo list'`,
	Run: func(cmd *cobra.Command, args []string) {
		list := cmd2.ListItem{}
		list.Add(args[0])
		fmt.Println(args[0] + " added to list")
	},
}

func init() {

	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
