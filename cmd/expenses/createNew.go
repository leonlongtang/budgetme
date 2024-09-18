/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package expenses

import (
	"fmt"

	"budgetme/sqldb"
	"budgetme/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ExpenseRaw struct {
	Amount   float64
	Category string
	Date     string
}

var (
	amount   float64
	category string
	date     string
	config   string
)

// createNewCmd represents the createNew command
var createNewCmd = &cobra.Command{
	Use:   "createNew",
	Short: "Adds a new expense in the database",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		logger := utils.GetLogger()

		// Check if config file is provided and load expenses from it
		if config != "" {
			logger.Info("Loading expenses from config file")
			var expenses []ExpenseRaw
			err := viper.UnmarshalKey("createNew.expenses", &expenses)
			if err != nil {
				logger.Error("Failed to load expenses from config: ", err)
				return
			}

			// Add each expense to the database
			for _, exp := range expenses {
				logger.Infof("Adding expense from config: %+v", exp)
				err := sqldb.AddExpense(db, exp.Amount, exp.Category, exp.Date)
				if err != nil {
					logger.Errorf("Failed to add expense from config: %+v, error: %v", exp, err)
				} else {
					fmt.Println("Added expense:", exp)
				}
			}
		} else {
			// Use flags for a single expense
			if amount == 0 || category == "" || date == "" {
				logger.Error("Missing required flags: amount, category, or date")
				fmt.Println("Error: missing required flags --amount, --category, and --date")
				return
			}

			logger.Infof("Adding expense from flags: amount=%f, category=%s, date=%s", amount, category, date)
			err := sqldb.AddExpense(db, amount, category, date)
			if err != nil {
				logger.Errorf("Failed to add expense from flags: %v", err)
			} else {
				fmt.Printf("Added expense: Amount=%.2f, Category=%s, Date=%s\n", amount, category, date)
			}
		}
	},
}

func init() {
	createNewCmd.Flags().StringVarP(&config, "config", "c", "", "Path to config file with expenses")

	createNewCmd.Flags().Float64VarP(&amount, "amount", "a", 0.0, "Expense amount")
	createNewCmd.Flags().StringVarP(&category, "category", "c", "", "Expense category")
	createNewCmd.Flags().StringVarP(&date, "date", "d", "", "Expense date (YYYY-MM-DD)")
	ExpensesCmd.AddCommand(createNewCmd)
}
