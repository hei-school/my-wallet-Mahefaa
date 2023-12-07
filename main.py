from core.Wallet import Wallet
from core.WalletRunner import WalletRunner

if __name__ == '__main__':
    wallet = Wallet()
    runner = WalletRunner(wallet)
    runner.run()