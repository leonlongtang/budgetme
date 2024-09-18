package expenses

import (
	"budgetme/models"
	"budgetme/sqldb" // Assuming you have sqldb for fetching expenses
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// categories slice to hold the passed categories from the user
var (
	format     string
	categories []string
)

// reportCmd represents the report command
var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Generate a report of expenses per category",
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize DB
		db, err := sqldb.InitDB()
		if err != nil {
			fmt.Println("Failed to initialize the database:", err)
			return
		}
		defer db.Close()

		// Fetch expenses from the database
		expenses, err := sqldb.FetchExpenses(db, "id", "asc")
		if err != nil {
			fmt.Println("Error fetching expenses:", err)
			return
		}

		outputTotalperCategory(expenses)

	},
}

func outputTotalperCategory(expenses []models.Expense) {
	categorySums := make(map[string]float64)
	if len(categories) > 0 {
		fmt.Println("Showing data for the following categories:", strings.Join(categories, ", "))
		for _, exp := range expenses {
			if containsCategory(exp.Category, categories) {
				categorySums[exp.Category] += exp.Amount
			}
		}
	} else {
		// If no categories are specified, sum all categories
		fmt.Println("Showing data for all categories")
		for _, exp := range expenses {
			categorySums[exp.Category] += exp.Amount
		}
	}

	// Print the result
	fmt.Printf("%-15s %-10s\n", "Category", "Total")
	fmt.Println("------------------------------")
	for category, total := range categorySums {
		fmt.Printf("%-15s %-10.2f\n", category, total)
	}
}

func containsCategory(category string, selectedCategories []string) bool {
	for _, c := range selectedCategories {
		if category == c {
			return true
		}
	}
	return false
}

func init() {
	// Add the --categories flag to the report command
	reportCmd.Flags().StringSliceVarP(&categories, "categories", "c", []string{}, "Filter by specific categories (comma-separated)")
	reportCmd.Flags().StringVarP(&format, "format", "f", "table", "Output format (table, json, csv)")

	// Add the report command to the expenses command
	ExpensesCmd.AddCommand(reportCmd)
}
