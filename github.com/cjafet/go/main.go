package main

import (
	"fmt"

	"github.com/cjafet/go/factory/payments"
)

func main() {
	payment, err := payments.GetPaymentMethod(3)
	if err != nil {
		fmt.Errorf("error is %v", err)
	}
	fmt.Print(payment.Pay(10))
}
