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

func (w *Wallet) GetBalance() (float64, error) {
	exception := w.checkActionValidity()
	if exception != nil {
		return 0, exception
	}
	fmt.Printf("Balance: $%.2f\n", w.balance)
	return w.balance, nil
}

func (w *Wallet) Withdraw(amount float64) (float64, error) {
	exception := w.checkActionValidity()
	if exception != nil {
		return 0, exception
	}

	if w.balance < amount {
		panic(errors.New("Error: Insufficient funds."))
	}

	w.balance -= amount
	fmt.Printf("Withdrawal of $%.2f successful. Remaining balance: $%.2f\n", amount, w.balance)
	return w.balance, nil
}

func (w *Wallet) StoreCard(cardName string) (string, error) {
	exception := w.checkActionValidity()
	if exception != nil {
		return "", exception
	}
	if _, exists := w.cards[cardName]; !exists {
		w.cards[cardName] = struct{}{}
		fmt.Printf("Card '%s' stored in the wallet.\n", cardName)
		return cardName, nil
	} else {
		return "", errors.New(fmt.Sprintf("Card already present, having two cards of the same name is not yet implemented."))
	}
}

func (w *Wallet) Deposit(amount float64) (float64, error) {
	exception := w.checkActionValidity()
	if exception != nil {
		return 0, exception
	}
	if amount < 0 {
		return 0, errors.New(fmt.Sprintf("Error: Invalid amount $%.2f", amount))
	}

	w.balance += amount
	fmt.Printf("Deposit of $%.2f successful. Remaining balance: $%.2f\n", amount, w.balance)
	return w.balance, nil
}

func (w *Wallet) GetCard(cardName string) (string, error) {
	exception := w.checkActionValidity()
	if exception != nil {
		return "", exception
	}
	if _, exists := w.cards[cardName]; exists {
		fmt.Printf("Card '%s' retrieved from the wallet.\n", cardName)
		return cardName, nil
	}
	return "", errors.New(fmt.Sprintf("Error: Card '%s' not found.", cardName))
}

func (w *Wallet) GetCards() (map[string]struct{}, error) {
	exception := w.checkActionValidity()
	return w.cards, exception
}

func (w *Wallet) checkActionValidity() error {
	if !w.isOpen {
		return errors.New("Error: Wallet is closed.")
	}
	return nil
}

func (w *Wallet) GetState() string {
	return fmt.Sprintf("Wallet {isOpen: %v}", w.isOpen)
}
