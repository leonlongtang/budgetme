package models

import "time"

type Expense struct {
	ID       int
	Amount   float64
	Category string
	Date     time.Time
}
