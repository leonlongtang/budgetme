package utils

import (
	"github.com/Rhymond/go-money"
)

func ConvertAmount(amount int64, base string, target string, rate float64) *money.Money {
	baseMoney := money.New(amount, base)
	targetAmount := baseMoney.Amount() * int64(rate*100) / 100
	targetMoney := money.New(targetAmount, target)
	return targetMoney
}
