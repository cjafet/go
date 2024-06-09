package payments

import (
	"fmt"
	"reflect"
)

type CashPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using %v\n", amount, reflect.TypeOf(c))
}
