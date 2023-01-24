/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Loa212/todo2/initializers"
	"github.com/Loa212/todo2/models"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [id] [?title] [?done]",
	Short: "update a todo",
	Long: ``,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		var todo models.Todo
		result := initializers.DB.Where("id = ?", id).First(&todo)

		if result.Error != nil {
			
			fmt.Printf("Todo %+v not found", id)
			
			return
		}

		if len(args) > 1 {
			todo.Title = args[1]
		}

		if len(args) > 2 {
			todo.Done = args[2] == "true"
		}

		result = initializers.DB.Save(&todo)

		if result.Error != nil {
			
			fmt.Printf("Can't update todo %+v", id)
			
			return
		}

		fmt.Printf("updated todo: %+v", id)

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
