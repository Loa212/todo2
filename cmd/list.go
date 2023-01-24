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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "this lists the todos",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {

		var todos []models.Todo

		result := initializers.DB.Find(&todos)

		if result.Error != nil {
			
			println("Error getting posts")
			
			return
		}
	
		fmt.Printf("%+v\n", todos)

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
