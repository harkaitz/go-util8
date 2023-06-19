package u

import (
	"github.com/Rhymond/go-money"
	"fmt"
)

// Money converts Money(3,60) to 360.
func Money(i,j money.Amount) (o money.Amount) {
	return i*100+j
}

// MoneyStr prints the amount in human form.
func MoneyStr(o money.Amount, c string, l string) (s string) {
	return fmt.Sprintf("%v.%02v€", o / 100, o % 100)
}

