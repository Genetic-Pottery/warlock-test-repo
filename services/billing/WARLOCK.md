<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# billing

The billing directory holds core billing domain types across multiple languages: invoices, ledgers, payments, and refunds, each with basic totaling or settlement logic and matching tests.

## Files

- `Invoice.java` (326 B) — Java class Invoice with constant MAX_LINES = 200 and total() summing 10 lines. · declares `Invoice`
- `InvoiceTest.java` (346 B) — Java test class InvoiceTest with test anInvoiceTotalsItsLines covering Invoice.
- `Ledger.kt` (291 B) — Kotlin class Ledger holding entries, sum() method, constant LEDGER_NAME, and function newLedger(). · declares `Ledger`, `fun`
- `LedgerTests.kt` (275 B) — Kotlin test class LedgerTests with test aLedgerSumsItsEntries covering Ledger.
- `Payment.cs` (269 B) — C# class Payment with constant RetryLimit = 5 and Settle(decimal amount) method. · declares `Payment`
- `Refund.swift` (347 B) — Swift struct Refund with amount field and isAllowed(), constant refundWindowDays = 30, and function makeRefund(amount:). · declares `Refund`, `isAllowed`, `makeRefund`

## Structure

- InvoiceTest.java exercises Invoice's total() method.
- LedgerTests.kt exercises Ledger's sum() method.
- newLedger() constructs a Ledger with no entries.
- makeRefund(amount:) constructs a Refund instance.

## Rules

- Invoice caps lines at MAX_LINES = 200.
- Payment limits retries via RetryLimit = 5.
- Refund enforces a refundWindowDays of 30.

## Where to look

- how many lines an invoice can hold → `Invoice.java` `MAX_LINES`
- how payment retries are limited → `Payment.cs` `RetryLimit`
- how long a refund can be requested after purchase → `Refund.swift` `refundWindowDays`
- how to create an empty ledger → `Ledger.kt` `newLedger`
- how to construct a refund → `Refund.swift` `makeRefund`
- tests for invoice totaling → `InvoiceTest.java` `anInvoiceTotalsItsLines`
- tests for ledger summing → `LedgerTests.kt` `aLedgerSumsItsEntries`
