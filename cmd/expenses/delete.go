/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package expenses

import (
	"budgetme/sqldb"
	"fmt"

	"github.com/spf13/cobra"
)

var id int

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an item from database",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := sqldb.DeleteExpense(db, id)
		if err != nil {
			fmt.Println("Error removing expense:", err)
			return
		}
		fmt.Printf("Expense ID == %d deleted successfully!\n", id)
	},
}

func init() {
	deleteCmd.Flags().IntVarP(&id, "id", "i", 0, "Expense ID")
	ExpensesCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
