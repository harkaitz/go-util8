package u

import (
	"github.com/Rhymond/go-money"
	"fmt"
)

func Money(i,j money.Amount) (o money.Amount) {
	return i*100+j
}

func MoneyStr(o money.Amount, c string, l string) (s string) {
	return fmt.Sprintf("%v.%02v€", o / 100, o % 100)
}

