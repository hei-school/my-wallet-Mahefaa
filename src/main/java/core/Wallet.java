package src.main.java.core;

import java.util.HashSet;
import java.util.Set;
import src.main.java.exception.CustomRuntimeException;

public class Wallet {
    private boolean isOpen;
    private double balance;
    private final Set<String> cards;

    public Wallet() {
        isOpen = false;
        balance = 0.0;
        cards = new HashSet<>();
    }

    public void open() {
        isOpen = true;
        System.out.println("Wallet opened.");
    }

    public void close() {
        isOpen = false;
        System.out.println("Wallet closed.");
    }

    public double getBalance() {
        checkActionValidity();
        System.out.println("Balance: $" + balance);
        return balance;
    }

    public double withdraw(double amount) {
        checkActionValidity();

        if (balance < amount) {
            throw new CustomRuntimeException("Error: Insufficient funds.");
        }

        balance -= amount;
        System.out.println("Withdrawal of $" + amount + " successful. Remaining balance: $" + balance);
        return balance;
    }

    public Set<String> storeCard(String cardName) {
        checkActionValidity();

        if (cards.add(cardName)) {
            System.out.println("Card '" + cardName + "' stored in the wallet.");
        } else {
            System.out.println("Card already present, having two cards of the same name is not yet implemented");
        }
        return cards;
    }

    public double deposit(double amount) {
        if (amount < 0) {
            throw new CustomRuntimeException("Error: Invalid amount ${amount}");
        }
        System.out.println("Deposit of $" + amount + " successful. Remaining balance: $" + this.balance);
        this.balance += amount;
        return this.balance;
    }

    public String getCard(String cardName) {
        checkActionValidity();

        if (cards.contains(cardName)) {
            System.out.println("Card '" + cardName + "' retrieved from the wallet.");
            return cardName;
        }
        throw new CustomRuntimeException("Error: Card '" + cardName + "' not found.");
    }

    public Set<String> getCards() {
        checkActionValidity();
        return cards;
    }

    private void checkActionValidity() {
        if (!isOpen) {
            throw new CustomRuntimeException("Error: Wallet is closed.");
        }
    }

    public String getState() {
        return "Wallet {isOpen: " + isOpen + "}";
    }
}
