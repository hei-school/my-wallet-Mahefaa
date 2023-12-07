import readline from 'readline';
import Wallet from "./core/Wallet.js";
import WalletApplicationRunner from "./core/WalletApplicationRunner.js";

const wallet = new Wallet();
const reader = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});
const runner = new WalletApplicationRunner(wallet, reader);
runner.run();