package payments

import (
	"strings"
	"testing"
)

func TestCreatePaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(1)
	if err != nil {
		t.Fatal("A payment method of type 'Cash' must exist")
	}

	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid using *payments.CashPM") {
		t.Error("The cash payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}
