package sqldb

import (
	"budgetme/models"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3" // Import your database driver
	"github.com/sirupsen/logrus"
)

// Database struct to hold DB connection and Logger
type Database struct {
	DB  *sql.DB
	Log *logrus.Logger
}

// InitDB initializes the database and returns a Database struct with the logger
func InitDB(log *logrus.Logger) (*Database, error) {
	log.Info("Initializing the database connection")
	db, err := sql.Open("sqlite3", "expenses.db")
	if err != nil {
		log.Error("Error opening database connection: ", err)
		return nil, err
	}

	// Return the Database struct
	return &Database{
		DB:  db,
		Log: log,
	}, nil
}

// AddExpense inserts a new expense into the database
func (d *Database) AddExpense(amount float64, category string, date string) error {
	d.Log.Info("Inserting a new expense into the database")

	insertQuery := `INSERT INTO expenses (amount, category, date) VALUES (?, ?, ?)`
	_, err := d.DB.Exec(insertQuery, amount, category, date)
	if err != nil {
		d.Log.Error("Error inserting expense: ", err)
		return err
	}

	d.Log.Info("Expense inserted successfully")
	return nil
}

func (d *Database) FetchExpenses(orderBy string, orderDirection string) ([]models.Expense, error) {
	validColumns := map[string]bool{"id": true, "amount": true, "category": true, "date": true}
	if !validColumns[orderBy] {
		d.Log.Errorf("Invalid orderBy column: %s", orderBy)
		return nil, fmt.Errorf("invalid orderBy column: %s", orderBy)
	}

	if orderDirection != "asc" && orderDirection != "desc" {
		d.Log.Errorf("Invalid orderDirection: %s", orderDirection)
		return nil, fmt.Errorf("invalid orderDirection: %s", orderDirection)
	}

	query := fmt.Sprintf(`SELECT id, amount, category, date FROM expenses ORDER BY %s %s`, orderBy, orderDirection)
	d.Log.Infof("Executing query: %s", query)

	rows, err := d.DB.Query(query)
	if err != nil {
		d.Log.Errorf("Error querying expenses: %v", err)
		return nil, err
	}
	defer rows.Close()

	var expenses []models.Expense
	for rows.Next() {
		var exp models.Expense
		var dateStr string

		err := rows.Scan(&exp.ID, &exp.Amount, &exp.Category, &dateStr)
		if err != nil {
			d.Log.Errorf("Error scanning row: %v", err)
			continue
		}

		exp.Date, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			d.Log.Errorf("Error parsing date: %v", err)
			continue
		}

		expenses = append(expenses, exp)
	}

	if err = rows.Err(); err != nil {
		d.Log.Errorf("Error iterating rows: %v", err)
		return nil, err
	}

	d.Log.Infof("Successfully fetched %d expenses", len(expenses))
	return expenses, nil
}

// DeleteExpense deletes an expense by ID and logs the operation
func (d *Database) DeleteExpense(id int) error {
	query := "DELETE FROM expenses WHERE id = ?"

	d.Log.Infof("Database layer: Deleting expense with ID %d", id)
	_, err := d.DB.Exec(query, id)
	if err != nil {
		d.Log.Errorf("Error deleting expense with ID %d: %v", id, err)
		return fmt.Errorf("error deleting expense: %v", err)
	}

	d.Log.Infof("Database layer: Successfully deleted expense with ID %d", id)
	return nil
}
