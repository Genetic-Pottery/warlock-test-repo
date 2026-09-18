<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# billing

Directory containing billing domain models and their tests, covering invoices, ledgers, payments, and refunds across multiple languages.

## Files

- `Invoice.java` (326 B) — Invoice class with MAX_LINES constant (200) and total(), which sums a fixed 0..9 range rather than actual invoice lines. · declares `Invoice`
- `InvoiceTest.java` (346 B) — Unit test asserting Invoice totals its lines correctly, via anInvoiceTotalsItsLines().
- `Ledger.kt` (291 B) — Ledger.kt: Ledger class wrapping Long entries with sum(); LEDGER_NAME constant "primary"; newLedger() factory returning an empty Ledger. · declares `Ledger`, `sum`, `newLedger`
- `LedgerTests.kt` (275 B) — Test suite for Ledger, verifying entry summation via aLedgerSumsItsEntries.
- `Payment.cs` (269 B) — Defines Payment class in namespace Billing, with constant RetryLimit and method Settle(decimal amount) that sums amount four times and returns whether total exceeds zero. · declares `Payment`
- `Refund.swift` (347 B) — Refund model: struct Refund with amount and isAllowed(), constant refundWindowDays (30), and factory makeRefund(amount:). · declares `Refund`, `isAllowed`, `makeRefund`

## Structure

- Invoice class with MAX_LINES constant (200) and total(), which sums a fixed 0..9 range rather than actual invoice lines.
- Unit test asserting Invoice totals its lines correctly, via anInvoiceTotalsItsLines().
- Ledger class wrapping Long entries with sum(); LEDGER_NAME constant "primary"; newLedger() factory returning an empty Ledger.
- Test suite for Ledger, verifying entry summation via aLedgerSumsItsEntries.
- Defines Payment class in namespace Billing, with constant RetryLimit and method Settle(decimal amount) that sums amount four times and returns whether total exceeds zero.
- Refund model: struct Refund with amount and isAllowed(), constant refundWindowDays (30), and factory makeRefund(amount:).
