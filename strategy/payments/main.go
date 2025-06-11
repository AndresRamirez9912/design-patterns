package main

import "fmt"

// Strategy interface which must be implemented by the payment methods
type IPayment interface {
	Pay(amount int)
}

// Context (Class / Obj which need to call the Pay method)
type Payment struct {
	paymentMethod IPayment
}

func (p Payment) PerformPay(amount int) {
	p.paymentMethod.Pay(amount)
}

// Implementations
type PaymentMethod func(int)

// paypal implementation
type Paypal PaymentMethod

func (p Paypal) Pay(amount int) {
	fmt.Printf("Paying %d with paypal \n", amount)
}

// PSE implementation
type PSE PaymentMethod

func (p PSE) Pay(amount int) {
	fmt.Printf("Paying %d with PSE \n", amount)
}

// Money implementation
type Money PaymentMethod

func (p Money) Pay(amount int) {
	fmt.Printf("Paying %d with Money \n", amount)
}

func main() {
	// create implementations
	var paypal Paypal

	var pse PSE

	var money Money

	// select payment method and create context
	var selectedMethod Payment
	method := "Paypal"
	switch method {
	case "Paypal":
		selectedMethod = Payment{paymentMethod: paypal}
	case "PSE":
		selectedMethod = Payment{paymentMethod: pse}
	case "Money":
		selectedMethod = Payment{paymentMethod: money}
	default:
		fmt.Println("Wrong payment method")
	}

	// call payment method
	selectedMethod.paymentMethod.Pay(12)
}
