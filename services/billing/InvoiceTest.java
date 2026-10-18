package billing;

import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertFalse;
import static org.junit.Assert.assertTrue;
import static org.junit.Assert.fail;

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

    @Test
    public void aNewInvoiceIsEmpty() {
        Invoice invoice = new Invoice(1L);
        assertTrue(invoice.isEmpty());
        assertEquals(0, invoice.lineCount());
    }

    @Test
    public void addLineCountsEachLine() {
        Invoice invoice = new Invoice(1L);
        invoice.addLine(500L);
        invoice.addLine(0L);
        assertFalse(invoice.isEmpty());
        assertEquals(2, invoice.lineCount());
    }

    @Test
    public void addLineRejectsANegativeAmount() {
        Invoice invoice = new Invoice(1L);
        try {
            invoice.addLine(-1L);
            fail("expected IllegalArgumentException");
        } catch (IllegalArgumentException expected) {
            // expected
        }
        assertTrue(invoice.isEmpty());
        assertEquals(0, invoice.lineCount());
    }

    @Test
    public void discountedByZeroKeepsTheTotal() {
        Invoice invoice = new Invoice(1L);
        assertEquals(45L, invoice.discounted(0));
    }

    @Test
    public void discountedByOneHundredIsZero() {
        Invoice invoice = new Invoice(1L);
        assertEquals(0L, invoice.discounted(100));
    }

    @Test
    public void discountedRoundsDown() {
        Invoice invoice = new Invoice(1L);
        // 45 * 90 / 100 = 40.5, which rounds down to 40.
        assertEquals(40L, invoice.discounted(10));
    }

    @Test(expected = IllegalArgumentException.class)
    public void discountedRejectsANegativePercent() {
        new Invoice(1L).discounted(-1);
    }

    @Test(expected = IllegalArgumentException.class)
    public void discountedRejectsAPercentAboveOneHundred() {
        new Invoice(1L).discounted(101);
    }
}
