package main

import "fmt"

type BankAccount struct {
	Owner   string
	Balance float64
}

func (account *BankAccount) Deposit(amount float64) {
	account.Balance += amount
}

func (account *BankAccount) Display() {
	fmt.Println("Balance: ", account.Balance)
}

func ApplyInterest(account *BankAccount, interest float64) {
	account.Balance = account.Balance + (interest/100)*account.Balance
}

func main() {
	newAccount := BankAccount{
		Owner:   "Tanmay Khanna",
		Balance: 10000,
	}
	ApplyInterest(&newAccount, 5)
	ApplyInterest(&newAccount, 10)
	//newAccount.Deposit(80000)
	newAccount.Display()

}
