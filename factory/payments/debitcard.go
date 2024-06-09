package payments

import (
	"fmt"
	"reflect"
)

type DebitCardPM struct{}

func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using %v\n", amount, reflect.TypeOf(c))
}
