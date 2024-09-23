package expenses

import (
	"budgetme/services"
	"fmt"
	"sort"

	"github.com/spf13/cobra"
)

var (
	groupBy string
)

// reportCmd represents the report command
var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Generate a report of expenses grouped by time intervals",
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize the service layer
		expenseService := services.NewExpenseService(db)

		// Fetch and group the expenses based on the groupBy flag
		grouped, err := expenseService.GenerateReport(groupBy)
		if err != nil {
			log.Error("Error generating report:", err)
			fmt.Println("Failed to generate report:", err)
			return
		}

		// Sort the intervals (year, month, or week)
		intervals := make([]string, 0, len(grouped))
		for interval := range grouped {
			intervals = append(intervals, interval)
		}
		sort.Strings(intervals)

		// Iterate through the sorted intervals and print the grouped expenses
		for _, interval := range intervals {
			exps := grouped[interval]
			var total float64

			fmt.Println("Time:", interval)
			fmt.Println("----------------------------------------------")
			expenseService.PrintExpenses(exps)

			// Calculate the total for the interval
			for _, exp := range exps {
				total += exp.Amount
			}

			fmt.Printf("Total: %.2f\n", total)
			fmt.Println()
		}
	},
}

func init() {
	reportCmd.Flags().StringVarP(&groupBy, "groupBy", "g", "year", "Group expenses by (year, month, week)")
	ExpensesCmd.AddCommand(reportCmd)
}
