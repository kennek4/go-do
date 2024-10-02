/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package subcmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

var taskName string

// taskCmd represents the task command
var TaskCmd = &cobra.Command{
	Use:   "task",
	Short: "Creates a task",
	Long: `

Creates a task that is placed under the current
working task group. 

The task can be placed into a different
task group if one is provided by using
the following flag:

	-g NAME_OF_TASK_GROUP

The arguments provided do not need to be surrounded 
in quotation marks (such as "" or '').

	`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("task called with arg: %s", args[0])
	},
}

func init() {
	TaskCmd.PersistentFlags().StringVarP(&taskName, "name", "n", "", "name of the task")
}
