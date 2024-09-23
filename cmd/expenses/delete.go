package expenses

import (
	"budgetme/services"
	"fmt"

	"github.com/spf13/cobra"
)

var id int

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an expense from the database",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize the service layer with the existing Database struct
		expenseService := services.NewExpenseService(db)

		// Call the service layer to delete the expense by ID
		err := expenseService.DeleteExpense(id)
		if err != nil {
			log.Error("Error deleting expense:", err)
			return
		}

		fmt.Printf("Expense with ID == %d deleted successfully!\n", id)
	},
}

func init() {
	deleteCmd.Flags().IntVarP(&id, "id", "i", 0, "Expense ID")
	ExpensesCmd.AddCommand(deleteCmd)
}
