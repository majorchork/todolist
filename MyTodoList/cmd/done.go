/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	cmd2 "github.com/majorchork/week4/mytodolist/MyTodoList/cmd/model"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "A brief description of your command",
	Long: `After a task, use this to mark a task as done
This is a soft delete function for the tasks you mark as done, only displaying undone tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		list := cmd2.ListItem{}
		list.Done(args[0])
		fmt.Println("item status updated")
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
