package src.main.java.core;

import static src.main.java.utils.ExceptionResolverUtils.tryCatchLogger;

import java.util.Scanner;

public class WalletApplicationRunner implements Runnable {
    private final Wallet wallet;
    private boolean isRunning;

    public static final String MENU_MESSAGE = """
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
    private final Scanner scanner;

    public WalletApplicationRunner(Wallet wallet, Scanner scanner) {
        assert (wallet != null) : " wallet must not be null for the runner to work well";
        assert (scanner != null) : "Scanner is invalid. User inputs are unreadable";
        this.wallet = wallet;
        this.scanner = scanner;
        isRunning = false;
    }

    private void showMenu() {
        System.out.println(wallet.getState());
        System.out.println(MENU_MESSAGE);
    }

    private void handleChoice(int choice) {
        switch (choice) {
            case 1 -> wallet.open();
            case 2 -> wallet.close();
            case 3 -> finallySkippedTryCatchLogger(wallet::getBalance);
            case 4 -> {
                System.out.print("Enter withdrawal amount: $");
                double withdrawalAmount = scanner.nextDouble();
                finallySkippedTryCatchLogger(() -> wallet.withdraw(withdrawalAmount));
            }
            case 5 -> {
                scanner.nextLine(); // Consume newline
                System.out.print("Enter card name to store: ");
                String cardName = scanner.nextLine();
                finallySkippedTryCatchLogger(() -> wallet.storeCard(cardName));
            }
            case 6 -> {
                scanner.nextLine(); // Consume newline
                System.out.print("Enter card name to retrieve: ");
                String retrieveCardName = scanner.nextLine();
                finallySkippedTryCatchLogger(() -> wallet.getCard(retrieveCardName));
            }
            case 7 -> finallySkippedTryCatchLogger(() -> System.out.println("cards = " + wallet.getCards()));
            case 8 -> {
                scanner.nextLine(); // Consume newline
                System.out.print("Enter amount to deposit: $");
                double depositAmount = scanner.nextDouble();
                finallySkippedTryCatchLogger(() -> wallet.deposit(depositAmount));
            }
            case 0 -> stop();
            default -> System.out.println("Invalid choice. Please enter a valid option.");
        }
    }

    @Override
    public void run() {
        do {
            isRunning = true;
            showMenu();

            int choice = scanner.nextInt();

            handleChoice(choice);
        } while (isRunning);
    }

    public void stop() {
        isRunning = false;
        System.out.println("Closing Wallet Application");
        closeResources();
    }

    public void closeResources() {
        scanner.close();
    }

    private void finallySkippedTryCatchLogger(Runnable tryCb) {
        tryCatchLogger(tryCb, () -> {
        });
    }
}