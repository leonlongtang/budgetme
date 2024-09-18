package utils

import (
	"budgetme/models"
	"fmt"
)

func PrintExpenses(expenses []models.Expense) {
	fmt.Printf("%-5s %-10s %-15s %-15s\n", "ID", "Amount", "Category", "Date")
	fmt.Println("----------------------------------------------")

	// Print each expense in a formatted way
	for _, exp := range expenses {
		fmt.Printf("%-5d %-10.2f %-15s %-15s\n", exp.ID, exp.Amount, exp.Category, exp.Date)
	}
}
