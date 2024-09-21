package expenses

import (
	"budgetme/models"
	"budgetme/sqldb" // Assuming you have sqldb for fetching expenses
	"budgetme/utils"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// categories slice to hold the passed categories from the user
var (
	groupBy string
	grouped map[string][]models.Expense
)

// reportCmd represents the report command
var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Generate a report of expenses per category",
	Run: func(cmd *cobra.Command, args []string) {
		// Fetch expenses from the database
		expenses, err := sqldb.FetchExpenses(db, "date", "asc")
		if err != nil {
			log.Error("Error fetching expenses:", err)
			return
		}

		switch groupBy {
		case "month":
			grouped = groupByMonth(expenses)
		case "week":
			grouped = groupByWeek(expenses)
		default:
			grouped = groupByYear(expenses)
		}

		intervals := make([]string, 0, len(grouped))

		for interval := range grouped {
			intervals = append(intervals, interval)
		}
		sort.Strings(intervals)

		// Iterate through the sorted intervals
		for _, interval := range intervals {
			exps := grouped[interval]
			var total float64

			fmt.Println("Time: ", interval)
			fmt.Println("----------------------------------------------")
			utils.PrintExpenses(exps)

			for _, exp := range exps {
				total += exp.Amount
			}

			fmt.Printf("Total: %.2f\n", total)
			fmt.Println()
		}
	},
}

func groupByYear(expenses []models.Expense) map[string][]models.Expense {
	grouped := make(map[string][]models.Expense)
	for _, expense := range expenses {
		year := strconv.Itoa(expense.Date.Year()) // Extract the year from Date
		grouped[year] = append(grouped[year], expense)
	}
	return grouped
}

// Group expenses by month (Year and Month)
func groupByMonth(expenses []models.Expense) map[string][]models.Expense {
	grouped := make(map[string][]models.Expense)
	for _, expense := range expenses {
		yearMonth := expense.Date.Format("2006-01") // Format the Date as "YYYY-MM"
		grouped[yearMonth] = append(grouped[yearMonth], expense)
	}
	return grouped
}

// Group expenses by week (Year and Week number)
func groupByWeek(expenses []models.Expense) map[string][]models.Expense {
	grouped := make(map[string][]models.Expense)

	for _, expense := range expenses {
		// Find the start of the week (Monday)
		startOfWeek := getStartOfWeek(expense.Date)

		// Calculate the end of the week (Sunday)
		endOfWeek := startOfWeek.AddDate(0, 0, 6) // Add 6 days to get the end of the week

		// Format as "YYYY-MM-DD - YYYY-MM-DD"
		yearWeek := fmt.Sprintf("%s to %s", startOfWeek.Format("2006-01-02"), endOfWeek.Format("2006-01-02"))

		// Group expenses by this week range
		grouped[yearWeek] = append(grouped[yearWeek], expense)
	}

	return grouped
}

func getStartOfWeek(date time.Time) time.Time {
	// Subtract the day of the week (Monday=1, ..., Sunday=7) minus 1 to get back to Monday
	offset := int(time.Monday - date.Weekday())
	if offset > 0 {
		offset = -6 // Adjust for when date is on Sunday (Weekday = 0)
	}
	startOfWeek := date.AddDate(0, 0, offset)
	return startOfWeek
}

func init() {
	// Add the --categories flag to the report command
	reportCmd.Flags().StringVarP(&groupBy, "groupBy", "g", "year", "Group expenses by (year, month, week)")

	// Add the report command to the expenses command
	ExpensesCmd.AddCommand(reportCmd)
}
