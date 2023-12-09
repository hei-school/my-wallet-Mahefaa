package core

import (
	"bufio"
	"fmt"
	"strconv"
	"wallet/exception"
)

type WalletApplicationRunner struct {
	wallet    *Wallet
	isRunning bool
	scanner   *bufio.Scanner
}

const menuMessage = `
Wallet Menu:
1. Open
2. Close
3. Get Balance
4. Withdraw
5. Store Card
6. Get Card
7. Get Cards
8. Deposit
0. Exit
Enter your choice:`

func NewWalletApplicationRunner(wallet *Wallet, scanner *bufio.Scanner) *WalletApplicationRunner {
	return &WalletApplicationRunner{
		wallet:    wallet,
		isRunning: false,
		scanner:   scanner,
	}
}

func (r *WalletApplicationRunner) showMenu() {
	fmt.Println(r.wallet.GetState())
	fmt.Println(menuMessage)
}

func (r *WalletApplicationRunner) handleChoice(choice int) {
	switch choice {
	case 1:
		r.wallet.Open()
	case 2:
		r.wallet.Close()
	case 3:
		result := finallySkippedResolveWithoutPanic(func() (any, error) {
			return r.wallet.GetBalance()
		})
		if result != nil {
			fmt.Printf("Balance: $%.2f\n", result)
		}
	case 4:
		fmt.Print("Enter withdrawal amount: $")
		withdrawalAmount, err := strconv.ParseFloat(r.scanner.Text(), 64)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			return
		}
		finallySkippedResolveWithoutPanic(func() (any, error) {
			return r.wallet.Withdraw(withdrawalAmount)
		})
	case 5:
		fmt.Print("Enter card name to store: ")
		cardName := r.scanner.Text()
		finallySkippedResolveWithoutPanic(func() (any, error) {
			return r.wallet.StoreCard(cardName)
		})
	case 6:
		fmt.Print("Enter card name to retrieve: ")
		retrieveCardName := r.scanner.Text()
		finallySkippedResolveWithoutPanic(func() (any, error) {
			return r.wallet.GetCard(retrieveCardName)
		})
	case 7:
		fmt.Println("cards =", finallySkippedResolveWithoutPanic(func() (any, error) {
			return r.wallet.GetCards()
		}))
	case 8:
		fmt.Print("Enter amount to deposit: $")
		depositAmount, err := strconv.ParseFloat(r.scanner.Text(), 64)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			return
		}
		finallySkippedResolveWithoutPanic(func() (any, error) {
			return r.wallet.Deposit(depositAmount)
		})
	case 0:
		r.stop()
	default:
		fmt.Println("Invalid choice. Please enter a valid option.")
	}
}

func (r *WalletApplicationRunner) Run() {
	for {
		r.isRunning = true
		r.showMenu()

		r.scanner.Scan()
		choice, err := strconv.Atoi(r.scanner.Text())
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		r.handleChoice(choice)
		if !r.isRunning {
			break
		}
	}
}

func (r *WalletApplicationRunner) stop() {
	r.isRunning = false
	fmt.Println("Closing Wallet Application")
}

func finallySkippedResolveWithoutPanic(tryFunc func() (any, error)) any {
	return exception.ResolveWithoutPanic(tryFunc, func() any {
		return nil
	})
}
