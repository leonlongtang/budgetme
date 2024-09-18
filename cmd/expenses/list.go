/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package expenses

import (
	"fmt"

	"budgetme/sqldb"
	"budgetme/utils"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the saved expenses",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		expenses, err := sqldb.FetchExpenses(db, orderBy, direction)
		if err != nil {
			fmt.Println("Error fetching expenses", err)
			return
		}
		utils.PrintExpenses(expenses)
	},
}

func init() {
	listCmd.Flags().StringVarP(&orderBy, "orderBy", "o", "id", "Value to order by")
	listCmd.Flags().StringVarP(&direction, "direction", "d", "asc", "Direction of order")
	ExpensesCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
