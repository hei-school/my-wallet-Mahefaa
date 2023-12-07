package core

import (
	"errors"
	"fmt"
)

type Wallet struct {
	isOpen  bool
	balance float64
	cards   map[string]struct{}
}

func NewWallet() *Wallet {
	return &Wallet{
		isOpen:  false,
		balance: 0.0,
		cards:   make(map[string]struct{}),
	}
}

func (w *Wallet) Open() {
	w.isOpen = true
	fmt.Println("Wallet opened.")
}

func (w *Wallet) Close() {
	w.isOpen = false
	fmt.Println("Wallet closed.")
}

func (w *Wallet) GetBalance() float64 {
	w.checkActionValidity()
	fmt.Printf("Balance: $%.2f\n", w.balance)
	return w.balance
}

func (w *Wallet) Withdraw(amount float64) float64 {
	w.checkActionValidity()

	if w.balance < amount {
		panic(errors.New("Error: Insufficient funds."))
	}

	w.balance -= amount
	fmt.Printf("Withdrawal of $%.2f successful. Remaining balance: $%.2f\n", amount, w.balance)
	return w.balance
}

func (w *Wallet) StoreCard(cardName string) {
	w.checkActionValidity()

	if _, exists := w.cards[cardName]; !exists {
		w.cards[cardName] = struct{}{}
		fmt.Printf("Card '%s' stored in the wallet.\n", cardName)
	} else {
		fmt.Println("Card already present, having two cards of the same name is not yet implemented.")
	}
}

func (w *Wallet) Deposit(amount float64) float64 {
	if amount < 0 {
		panic(errors.New(fmt.Sprintf("Error: Invalid amount $%.2f", amount)))
	}

	w.balance += amount
	fmt.Printf("Deposit of $%.2f successful. Remaining balance: $%.2f\n", amount, w.balance)
	return w.balance
}

func (w *Wallet) GetCard(cardName string) string {
	w.checkActionValidity()

	if _, exists := w.cards[cardName]; exists {
		fmt.Printf("Card '%s' retrieved from the wallet.\n", cardName)
		return cardName
	}
	panic(errors.New(fmt.Sprintf("Error: Card '%s' not found.", cardName)))
}

func (w *Wallet) GetCards() map[string]struct{} {
	w.checkActionValidity()
	return w.cards
}

func (w *Wallet) checkActionValidity() {
	if !w.isOpen {
		panic(errors.New("Error: Wallet is closed."))
	}
}

func (w *Wallet) GetState() string {
	return fmt.Sprintf("Wallet {isOpen: %v}", w.isOpen)
}
