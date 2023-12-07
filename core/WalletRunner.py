from core.Wallet import Wallet
from utils.exception_resolver import try_except_logger

MENU_MESSAGE = """
            Wallet Menu:
            1. Open
            2. Close
            3. Get Balance
            4. Withdraw
            5. Store Card
            6. Get Card
            7. Get cards
            8. Deposit
            0. Exit
            Enter your choice:""";


def finally_skipped_try_except_logger(try_cb):
    try_except_logger(try_cb, lambda: {})


class WalletRunner:
    def __init__(self, wallet: Wallet):
        self.wallet = wallet
        self.is_running = False

    def show_menu(self):
        print(self.wallet.get_state())
        print(MENU_MESSAGE)

    def handle_choice(self, choice: int):
        match choice:
            case 1:
                self.wallet.open()
            case 2:
                self.wallet.close()
            case 3:
                try_except_logger(lambda: self.wallet.get_balance(), lambda: {})
            case 4:
                withdrawal_amount = float(input("Enter withdrawal amount: $"))
                finally_skipped_try_except_logger(lambda: self.wallet.withdraw(withdrawal_amount))
            case 5:
                card_name = input("Enter card name to store: ")
                finally_skipped_try_except_logger(lambda: self.wallet.store_card(card_name))
            case 6:
                card_name = input("Enter card name to retrieve: ")
                finally_skipped_try_except_logger(lambda: self.wallet.get_card(card_name))
            case 7:
                finally_skipped_try_except_logger(lambda: print(self.wallet.get_cards()))
            case 8:
                deposit_amount = float(input("Enter amount to deposit: $"))
                finally_skipped_try_except_logger(lambda: self.wallet.deposit(deposit_amount))
            case 0:
                self.stop()

    def stop(self):
        self.is_running = False
        print("Closing Wallet Application")

    def run(self):
        self.is_running = True

        while self.is_running:
            self.show_menu()

            choice = int(input(""))

            self.handle_choice(choice)

    def keep_going(self):
        if self.is_running:
            self.run()
        else:
            print("Wallet application closed")
            exit(0)
