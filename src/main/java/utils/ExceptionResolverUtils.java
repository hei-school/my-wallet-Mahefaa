package src.main.java.utils;

import src.main.java.exception.CustomRuntimeException;

public class ExceptionResolverUtils {
    public static void tryCatchLogger(Runnable tryCb, Runnable finallyCb) {
        try {
            tryCb.run();
        } catch (CustomRuntimeException e) {
            System.err.println(e.getMessage());
        }finally {
            finallyCb.run();
        }
    }
}
