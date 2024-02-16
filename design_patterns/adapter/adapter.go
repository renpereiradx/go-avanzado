package main

import "fmt"

type Payment interface {
	Pay()
}

type CashPayment struct{}

func (CashPayment) Pay() {
	fmt.Println("Payment using Cash")
}

type BankPayment struct{}

func (BankPayment) Pay(account int) {
	fmt.Printf("Paying using Bank Account: %d\n", account)
}

type BankPaymentAdapter struct {
	BankPayment *BankPayment
	bankAccount int
}

func (bpa BankPaymentAdapter) Pay() {
	bpa.BankPayment.Pay(bpa.bankAccount)
}

func ProcessPayment(p Payment) {
	p.Pay()
}

func main() {
	cashPayment := &CashPayment{}
	ProcessPayment(cashPayment)
	bpa := &BankPaymentAdapter{
		BankPayment: &BankPayment{},
		bankAccount: 5,
	}
	ProcessPayment(bpa)
}
