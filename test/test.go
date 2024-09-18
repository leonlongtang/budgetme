package main

import (
	"budgetme/utils"
	"fmt"
	"os"
)

var (
	base   = "EUR"
	target = "USD"
	amount = 184
)

func main() {
	apiKey := os.Getenv("FIXER_API_KEY")
	if apiKey == "" {
		fmt.Println("FIXER_API_KEY is not set in the environment")
		return
	}

	rate, err := utils.GetConversionRate(apiKey, target)
	if err != nil {
		fmt.Println(err)
	}

	target_amount := utils.ConvertAmount(int64(amount), base, target, rate)

	fmt.Println(target_amount)
}
