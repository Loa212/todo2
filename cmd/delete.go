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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: ``,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		id := args[0]

		var todo models.Todo
		result := initializers.DB.Where("id = ?", id).First(&todo)

		if result.Error != nil {
			
			fmt.Printf("Todo %+v not found", id)
			
			return
		}

		result = initializers.DB.Delete(&todo)

		if result.Error != nil {
			
			fmt.Printf("Can't delete todo %+v", id)
			
			return
		}

	
		fmt.Printf("deleted todo: %+v", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
