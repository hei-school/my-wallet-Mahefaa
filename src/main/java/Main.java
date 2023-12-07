package src.main.java;

import src.main.java.core.Wallet;
import src.main.java.core.WalletApplicationRunner;
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Wallet wallet = new Wallet();
        try(Scanner scanner = new Scanner(System.in)) {
            Runnable walletApplicationRunner = new WalletApplicationRunner(wallet, scanner);
            walletApplicationRunner.run();
        }
    }
}
