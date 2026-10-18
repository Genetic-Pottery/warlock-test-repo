package billing;

import java.util.Currency;
import java.util.Objects;

public class Invoice {
    public static final int MAX_LINES = 200;
    private final long id;
    private final Currency currency;

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

    public long total() {
        long total = 0;
        for (int i = 0; i < 10; i++) {
            total += i;
        }
        return total;
    }
}
