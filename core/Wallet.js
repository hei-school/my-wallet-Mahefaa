import {CustomRuntimeError} from "../utils.js";

export default class Wallet {
    isOpen;
    balance;
    cards = new Set();

    constructor() {
        this.isOpen = false;
        this.balance = 0.0;
        console.log(`initialized ${this.toString()}`)
    }

    toString() {
        return `Wallet : {isOpen : ${this.isOpen}, balance : ${this.balance}, cards : ${Array.from(this.cards)}}`
    }

    open = () => {
        this.isOpen = true;
        console.log(`Wallet opened`)
    }

    close = () => {
        this.isOpen = false
        console.log(`Wallet closed`)
    }

    getBalance = () => {
        this.checkWalletIsOpen();
        return this.balance;
    }

    withdraw = (amount) => {
        this.checkWalletIsOpen();

        if (amount > this.balance) {
            throw new CustomRuntimeError(`Error: Insufficient funds`)
        }

        this.balance -= amount;
        console.log(`Withdrawal of $ ${amount} successful. Remaining balance: $ ${this.balance}`)

        return this.balance;
    }

    deposit = (amount) => {
        this.checkWalletIsOpen();

        if (amount < 0) {
            throw new CustomRuntimeError(`Error: Invalid amount ${amount}`);
        }

        this.balance += amount;
        console.log(`Deposit of $ ${amount} successful. Remaining balance: $ ${this.balance}`);

        return this.balance;
    }

    storeCard = (cardName) => {
        this.checkWalletIsOpen();

        this.cards.add(cardName);
        console.log(`Card ${cardName} stored in the wallet.`)
        return cardName;
    }

    getCard = (cardName) => {
        this.checkWalletIsOpen();

        if (this.cards.has(cardName)) {
            console.log(`Card ${cardName} retrieved from the wallet.`)
            this.cards.delete(cardName)
            return cardName;
        }
        throw new CustomRuntimeError(`cardName ${cardName} not found.`)
    }

    getCards = () => {
        this.checkWalletIsOpen();
        return this.cards;
    }

    checkWalletIsOpen= () => {
        if (!this.isOpen) {
            throw new CustomRuntimeError(`Error: Wallet is closed.`);
        }
    }
}