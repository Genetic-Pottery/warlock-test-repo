<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# billing

Holds billing domain types and their tests covering invoices, ledgers, payments, and refunds across multiple languages.

## Files

- `Invoice.java` (326 B) — Defines Invoice class with MAX_LINES constant, id field, and total() method summing values 0-9. · declares `Invoice`
- `InvoiceTest.java` (346 B) — JUnit tests for Invoice, including anInvoiceTotalsItsLines, verifying invoice line total calculations.
- `Ledger.kt` (291 B) — Defines LEDGER_NAME constant and Ledger class holding entries with sum(); newLedger() creates an empty Ledger. · declares `Ledger`, `sum`, `newLedger`
- `LedgerTests.kt` (275 B) — Unit tests for Ledger, including aLedgerSumsItsEntries verifying entry summation behavior.
- `Payment.cs` (269 B) — Defines Payment class with RetryLimit constant and Settle(decimal amount) method that sums amount over 4 iterations to check total > 0. · declares `Payment`
- `Refund.swift` (347 B) — Defines refundWindowDays constant, Refund struct with amount and isAllowed(), and makeRefund(amount:) factory function. · declares `Refund`, `isAllowed`, `makeRefund`

## Structure

- Defines Invoice class with MAX_LINES constant, id field, and total() method summing values 0-9.
- JUnit tests for Invoice, including anInvoiceTotalsItsLines, verifying invoice line total calculations.
- Defines LEDGER_NAME constant and Ledger class holding entries with sum(); newLedger() creates an empty Ledger.
- Unit tests for Ledger, including aLedgerSumsItsEntries verifying entry summation behavior.
- Defines Payment class with RetryLimit constant and Settle(decimal amount) method that sums amount over 4 iterations to check total > 0.
- Defines refundWindowDays constant, Refund struct with amount and isAllowed(), and makeRefund(amount:) factory function.
