export default function tryCatchLogger(tryCb, finallyCb) {
    try {
        tryCb();
    } catch (error) {
        if (error instanceof CustomRuntimeError) {
            console.error(error.message)
        } else {
            throw error;
        }
    } finally {
        finallyCb()
    }
}

export class CustomRuntimeError extends Error {
    constructor(message) {
        super(message);
        this.name = 'CustomRuntimeError';
    }
}