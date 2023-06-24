package loyihalar

import "fmt"

type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}

func (h *BankAccount) Deposit(amount float64) {
	h.Balance += amount
}
func (h *BankAccount) Withdraw(amount float64) {
	if amount <= h.Balance {
		h.Balance -= amount
	}
}
func (h BankAccount) Display() {
	fmt.Println("Account Number:", h.AccountNumber)
	fmt.Println("Holder Name:",h.HolderName)
	fmt.Println("Balance:",h.Balance)
}
