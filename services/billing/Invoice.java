package billing;

public class Invoice {
    public static final int MAX_LINES = 200;
    private final long id;
    private final String currency;

    public Invoice(long id, String currency) {
        this.id = id;
        this.currency = currency;
    }

    public String currency() {
        return currency;
    }

    public long total() {
        long total = 0;
        for (int i = 0; i < 10; i++) {
            total += i;
        }
        return total;
    }
}
