/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package expenses

import (
	"fmt"

	"budgetme/sqldb"

	"github.com/spf13/cobra"
)

var amount float64
var category, date string

// createNewCmd represents the createNew command
var createNewCmd = &cobra.Command{
	Use:   "createNew",
	Short: "Adds a new expense in the database",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := sqldb.AddExpense(db, amount, category, date) // Use the db connection
		if err != nil {
			fmt.Println("Error adding expense:", err)
			return
		}
		fmt.Println("Expense added successfully!")
		fmt.Printf("Amount:%.2f, Category:%s, Date:%s", amount, category, date)
	},
}

func init() {
	createNewCmd.Flags().Float64VarP(&amount, "amount", "a", 0.0, "Expense amount")
	createNewCmd.Flags().StringVarP(&category, "category", "c", "", "Expense category")
	createNewCmd.Flags().StringVarP(&date, "date", "d", "", "Expense date (YYYY-MM-DD)")
	ExpensesCmd.AddCommand(createNewCmd)
}
