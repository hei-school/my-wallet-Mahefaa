from exception.CustomException import CustomException


class Wallet:
    def __init__(self):
        self.is_open = False
        self.balance = 0.0
        self.cards = set()

    def open(self):
        self.is_open = True

    def close(self):
        self.is_open = False

    def get_balance(self):
        self.check_action_validity()
        print(f"Balance: $ {self.balance}")
        return

    def withdraw(self, amount: float):
        self.check_action_validity()
        if amount > self.balance:
            raise CustomException("Error: Insufficient funds.")
        self.balance = self.balance - amount
        print(f"Withdrawal of $ {amount} successful. Remaining balance: $ {self.balance}")

    def store_card(self, card_name):
        self.check_action_validity()
        if card_name in self.cards:
            print(f"Card already present, having two cards of the same name is not yet implemented")
        self.cards.add(card_name)
        print(f"Card ${card_name} stored in the wallet.")
        return self.cards

    def deposit(self, amount: float):
        self.check_action_validity()
        if amount < 0:
            raise CustomException(f"Error: Invalid amount ${amount}")
        self.balance = self.balance + amount
        print(f"Deposit of $ {amount} successful. Remaining balance: ${self.balance}")

    def get_card(self, card_name):
        self.check_action_validity()
        if card_name in self.cards:
            print(f"Card {card_name} retrieved from the wallet.")
            self.cards.remove(card_name)
            return card_name
        raise CustomException(f"Error: Card {card_name} not found.")

    def get_cards(self):
        self.check_action_validity()
        return self.cards

    def get_state(self):
        return f"Wallet {{isOpen: {self.is_open}}}"

    def check_action_validity(self):
        if not self.is_open:
            raise CustomException("Error: Wallet is closed.")
