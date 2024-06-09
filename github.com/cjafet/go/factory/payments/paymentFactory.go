package payments

import (
	"fmt"
)

const (
	Cash       = 1
	DebitCard  = 2
	CreditCard = 3
)

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	case CreditCard:
		return new(CreditCardPM), nil
	default:
		return nil, fmt.Errorf(fmt.Sprintf("Payment method %d not recognized\n", m))
	}
}
