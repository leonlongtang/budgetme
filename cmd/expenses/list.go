package expenses

import (
	"budgetme/services"
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the saved expenses",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize the service layer
		expenseService := services.NewExpenseService(db)

		// Fetch the expenses
		expenses, err := expenseService.FetchExpenses(orderBy, direction)
		if err != nil {
			log.Error("Error fetching expenses:", err)
			fmt.Println("Failed to fetch expenses:", err)
			return
		}

		// Print the expenses using a utility function
		expenseService.PrintExpenses(expenses)
	},
}

func init() {
	listCmd.Flags().StringVarP(&orderBy, "orderBy", "o", "id", "Column to order by")
	listCmd.Flags().StringVarP(&direction, "direction", "d", "asc", "Direction of ordering")
	ExpensesCmd.AddCommand(listCmd)
}
