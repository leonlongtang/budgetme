package services

import (
	"budgetme/models"
	"budgetme/sqldb"
	"fmt"
	"strconv"
	"time"
)

// ExpenseService is the service layer for handling business logic related to expenses
type ExpenseService struct {
	db *sqldb.Database // The Database struct holding DB connection and logger
}

// NewExpenseService initializes the service with the database
func NewExpenseService(db *sqldb.Database) *ExpenseService {
	return &ExpenseService{db: db}
}

// AddExpense handles the logic for adding a new expense
func (s *ExpenseService) AddExpense(amount float64, category, date string) error {
	// You can add any additional business logic here before saving
	s.db.Log.Info("Service layer: Adding a new expense")

	// Call the database method to insert the expense
	err := s.db.AddExpense(amount, category, date)
	if err != nil {
		s.db.Log.Error("Service layer: Failed to add expense", err)
		return fmt.Errorf("could not add expense: %v", err)
	}

	s.db.Log.Info("Service layer: Expense added successfully")
	return nil
}

// FetchExpenses handles retrieving expenses from the database with ordering
func (s *ExpenseService) FetchExpenses(orderBy, orderDirection string) ([]models.Expense, error) {
	s.db.Log.Info("Service layer: Fetching expenses")

	// Validate and process input, then call the database method
	expenses, err := s.db.FetchExpenses(orderBy, orderDirection)
	if err != nil {
		s.db.Log.Error("Service layer: Failed to fetch expenses", err)
		return nil, fmt.Errorf("could not fetch expenses: %v", err)
	}

	s.db.Log.Infof("Service layer: Successfully fetched %d expenses", len(expenses))
	return expenses, nil
}

// DeleteExpense handles deleting an expense by ID
func (s *ExpenseService) DeleteExpense(id int) error {
	s.db.Log.Infof("Service layer: Deleting expense with ID %d", id)

	// Call the database method to delete the expense
	err := s.db.DeleteExpense(id)
	if err != nil {
		s.db.Log.Error("Service layer: Failed to delete expense", err)
		return fmt.Errorf("could not delete expense with ID %d: %v", id, err)
	}

	s.db.Log.Infof("Service layer: Successfully deleted expense with ID %d", id)
	return nil
}

// GenerateReport fetches expenses and groups them by the specified interval (year, month, week)
func (s *ExpenseService) GenerateReport(groupBy string) (map[string][]models.Expense, error) {
	s.db.Log.Infof("Service layer: Generating report grouped by %s", groupBy)

	// Fetch all expenses from the database
	expenses, err := s.db.FetchExpenses("date", "asc")
	if err != nil {
		s.db.Log.Errorf("Service layer: Error fetching expenses for report: %v", err)
		return nil, fmt.Errorf("could not fetch expenses for report: %v", err)
	}

	// Group the expenses based on the specified groupBy value
	switch groupBy {
	case "month":
		return s.groupByMonth(expenses), nil
	case "week":
		return s.groupByWeek(expenses), nil
	default:
		return s.groupByYear(expenses), nil
	}
}

// Group expenses by year
func (s *ExpenseService) groupByYear(expenses []models.Expense) map[string][]models.Expense {
	grouped := make(map[string][]models.Expense)
	for _, expense := range expenses {
		year := strconv.Itoa(expense.Date.Year()) // Extract the year from Date
		grouped[year] = append(grouped[year], expense)
	}
	return grouped
}

// Group expenses by month (Year and Month)
func (s *ExpenseService) groupByMonth(expenses []models.Expense) map[string][]models.Expense {
	grouped := make(map[string][]models.Expense)
	for _, expense := range expenses {
		yearMonth := expense.Date.Format("2006-01") // Format the Date as "YYYY-MM"
		grouped[yearMonth] = append(grouped[yearMonth], expense)
	}
	return grouped
}

// Group expenses by week (Year and Week number)
func (s *ExpenseService) groupByWeek(expenses []models.Expense) map[string][]models.Expense {
	grouped := make(map[string][]models.Expense)

	for _, expense := range expenses {
		// Find the start of the week (Monday)
		startOfWeek := s.getStartOfWeek(expense.Date)

		// Calculate the end of the week (Sunday)
		endOfWeek := startOfWeek.AddDate(0, 0, 6) // Add 6 days to get the end of the week

		// Format as "YYYY-MM-DD - YYYY-MM-DD"
		yearWeek := fmt.Sprintf("%s to %s", startOfWeek.Format("2006-01-02"), endOfWeek.Format("2006-01-02"))

		// Group expenses by this week range
		grouped[yearWeek] = append(grouped[yearWeek], expense)
	}

	return grouped
}

// Helper function to get the start of the week (Monday)
func (s *ExpenseService) getStartOfWeek(date time.Time) time.Time {
	// Subtract the day of the week (Monday=1, ..., Sunday=7) minus 1 to get back to Monday
	offset := int(time.Monday - date.Weekday())
	if offset > 0 {
		offset = -6 // Adjust for when date is on Sunday (Weekday = 0)
	}
	return date.AddDate(0, 0, offset)
}

func (s *ExpenseService) PrintExpenses(expenses []models.Expense) {
	fmt.Printf("%-5s %-10s %-15s %-15s\n", "ID", "Amount", "Category", "Date")
	fmt.Println("----------------------------------------------")

	// Print each expense in a formatted way
	for _, exp := range expenses {
		fmt.Printf("%-5d %-10.2f %-15s %-15s\n", exp.ID, exp.Amount, exp.Category, exp.Date.Format("2006-01-02"))
	}
}
