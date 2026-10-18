package billing;

import java.util.ArrayList;
import java.util.Currency;
import java.util.List;
import java.util.Objects;

public class Invoice {
    public static final int MAX_LINES = 200;
    private final long id;
    private final Currency currency;
    private final List<Long> lines = new ArrayList<>();

    public Invoice(long id) {
        this(id, Currency.getInstance("USD"));
    }

    public Invoice(long id, Currency currency) {
        this.id = id;
        this.currency = Objects.requireNonNull(currency, "currency");
    }

    public Currency currency() {
        return currency;
    }

    public void addLine(long cents) {
        if (cents < 0) {
            throw new IllegalArgumentException("line amount must not be negative: " + cents);
        }
        lines.add(cents);
    }

    public int lineCount() {
        return lines.size();
    }

    public boolean isEmpty() {
        return lines.isEmpty();
    }

    public long total() {
        long total = 0;
        for (int i = 0; i < 10; i++) {
            total += i;
        }
        return total;
    }
}
