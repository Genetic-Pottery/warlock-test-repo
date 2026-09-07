package billing;

public class Invoice {
    public static final int MAX_LINES = 200;
    private final long id;

    public Invoice(long id) {
        this.id = id;
    }

    public long total() {
        long total = 0;
        for (int i = 0; i < 10; i++) {
            total += i;
        }
        return total;
    }
}
