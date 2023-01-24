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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [title] [done?]",
	Short: "Adds a todo. Pass a boolean value as a second arg to mark it as completed upon creation (defaults at completed=false).",
	Long: ``,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		


		//create post
	todo := models.Todo{Title: args[0], Done: (args[1] == "true"),}

	result := initializers.DB.Create(&todo)

	if result.Error != nil {
			
		println("Error creating post")
		
		return
	}

	fmt.Printf("%+v\n", todo)
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
