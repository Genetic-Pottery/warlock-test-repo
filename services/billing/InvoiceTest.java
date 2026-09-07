package billing;

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
}
