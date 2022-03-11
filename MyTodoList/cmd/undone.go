/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	cmd2 "github.com/majorchork/week4/mytodolist/MyTodoList/cmd/model"
	"github.com/spf13/cobra"
)

// undoneCmd represents the undone command
var undoneCmd = &cobra.Command{
	Use:   "undone",
	Short: "A brief description of your command",
	Long: `A task that was not done but mistakenly marked as done can be marked undone back with this command
It is used to retrieve a task`,
	Run: func(cmd *cobra.Command, args []string) {
		list := cmd2.ListItem{}
		list.UnDone(args[0])
		fmt.Println("undone called")
	},
}

func init() {
	rootCmd.AddCommand(undoneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// undoneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// undoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
