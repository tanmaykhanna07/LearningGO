package main

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float64)
}

type CreditCard struct {
	CardNumber string
}
type Bitcoin struct {
	WalletAddress string
}

func (b Bitcoin) Pay(amount float64) {
	fmt.Println("Paid ", amount, " using Bitcoin: ", b.WalletAddress)
}

func (c CreditCard) Pay(amount float64) {
	fmt.Println("Paid ", amount, " using Credit Card: ", c.CardNumber)
}

func ProcessPayment(method PaymentMethod, amount float64) error {
	if amount <= 0 {
		return errors.New("amount cannot be 0 or less than 0")
	}
	method.Pay(amount)
	return nil
}

func main() {
	card := CreditCard{CardNumber: "9801879107"}
	coin := Bitcoin{WalletAddress: "meowy@walletapp.com"}
	err1 := ProcessPayment(card, 9000)
	if err1 != nil {
		fmt.Println("Something's not right: ", err1)
		return
	}
	err2 := ProcessPayment(coin, -500)
	if err2 != nil {
		fmt.Println("Something's not right: ", err2)
		return
	}
}
