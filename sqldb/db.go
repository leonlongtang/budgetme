package sqldb

import (
	"database/sql"
	"fmt"

	"budgetme/models"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	// Initialize the SQLite database
	db, err := sql.Open("sqlite3", "./expenses.db")
	if err != nil {
		return nil, err
	}

	// Create the expenses table if it doesn't exist
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS expenses (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		amount REAL,
		category TEXT,
		date TEXT
	);`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, fmt.Errorf("error creating table: %v", err)
	}

	return db, nil
}

// Example function to add an expense
func AddExpense(db *sql.DB, amount float64, category string, date string) error {
	insertQuery := `INSERT INTO expenses (amount, category, date) VALUES (?, ?, ?)`
	_, err := db.Exec(insertQuery, amount, category, date)
	if err != nil {
		return fmt.Errorf("error inserting expense: %v", err)
	}
	return nil
}

func FetchExpenses(db *sql.DB, orderBy string, orderDirection string) ([]models.Expense, error) {
	validColumns := map[string]bool{"id": true, "amount": true, "category": true, "date": true}
	if !validColumns[orderBy] {
		return nil, fmt.Errorf("invalid orderBy column: %s", orderBy)
	}

	if orderDirection != "asc" && orderDirection != "desc" {
		return nil, fmt.Errorf("invalid orderDirectionL %s", orderDirection)
	}

	query := fmt.Sprintf(`SELECT id, amount, category, date FROM expenses ORDER BY %s %s`, orderBy, orderDirection)

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying expenses: %v", err)
	}

	defer rows.Close()

	var expenses []models.Expense

	for rows.Next() {
		var exp models.Expense

		err := rows.Scan(&exp.ID, &exp.Amount, &exp.Category, &exp.Date)
		if err != nil {
			fmt.Printf("Error scanning row: %v", err)
			continue
		}

		expenses = append(expenses, exp)
	}

	err = rows.Err()

	if err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return expenses, nil
}

func DeleteExpense(db *sql.DB, id int) error {
	query := "DELETE FROM expenses WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
