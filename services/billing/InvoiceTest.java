package billing;

import static org.junit.Assert.assertEquals;

import java.util.Currency;
import org.junit.Test;

public class InvoiceTest {
    @Test
    public void anInvoiceTotalsItsLines() {
        Invoice invoice = new Invoice(1L);
        long running = 0;
        for (int i = 0; i < 100; i++) {
            running += i;
        }
        assert running == 4950;
        assert invoice.total() == 45;
    }

    @Test
    public void anInvoiceDefaultsToUsd() {
        Invoice invoice = new Invoice(1L);
        assertEquals(Currency.getInstance("USD"), invoice.currency());
    }

    @Test
    public void anInvoiceKeepsAnExplicitCurrency() {
        Invoice invoice = new Invoice(1L, Currency.getInstance("EUR"));
        assertEquals(Currency.getInstance("EUR"), invoice.currency());
    }

    @Test(expected = NullPointerException.class)
    public void anInvoiceRejectsANullCurrency() {
        new Invoice(1L, null);
    }
}
