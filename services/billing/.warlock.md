<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# billing

Billing-related types across multiple languages (Java, Kotlin, C#, Swift) modeling invoices, ledgers, payments, and refunds, each with associated summing or validation logic.

## Files

- `Invoice.java` (326 B) — Defines Invoice class with MAX_LINES constant and total() method summing 0..9; constructed with an id, id unused in total(). · declares `Invoice`
- `InvoiceTest.java` (346 B) — JUnit test class InvoiceTest with anInvoiceTotalsItsLines(), verifying Invoice sums its line items into a total.
- `Ledger.kt` (291 B) — Ledger.kt: Ledger class wrapping a List<Long> of entries with sum(); newLedger() factory returns an empty Ledger; LEDGER_NAME const "primary". · declares `Ledger`, `sum`, `newLedger`
- `LedgerTests.kt` (275 B) — Unit tests for a Ledger type, verifying it sums its entries correctly.
- `Payment.cs` (269 B) — Payment class with RetryLimit=5 constant and Settle(decimal amount) method summing 4x amount to check total > 0. · declares `Payment`
- `Refund.swift` (347 B) — Defines Refund struct (amount) with isAllowed(), refundWindowDays constant (30), and makeRefund(amount:) factory function. · declares `Refund`, `isAllowed`, `makeRefund`

## Structure

- Defines Invoice class with MAX_LINES constant and total() method summing 0..9; constructed with an id, id unused in total().
- JUnit test class InvoiceTest with anInvoiceTotalsItsLines(), verifying Invoice sums its line items into a total.
- Ledger class wrapping a List<Long> of entries with sum(); newLedger() factory returns an empty Ledger; LEDGER_NAME const "primary".
- Unit tests for a Ledger type, verifying it sums its entries correctly.
- Payment class with RetryLimit=5 constant and Settle(decimal amount) method summing 4x amount to check total > 0.
- Defines Refund struct (amount) with isAllowed(), refundWindowDays constant (30), and makeRefund(amount:) factory function.
