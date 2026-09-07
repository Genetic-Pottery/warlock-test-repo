<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# billing

The billing directory holds core billing domain types for invoices, ledgers, payments, and refunds, each implemented in a different language with matching tests.

## Files

- `Invoice.java` (326 B) — Java Invoice class with MAX_LINES constant and total() method summing line entries. · declares `Invoice`
- `InvoiceTest.java` (346 B) — JUnit tests for Invoice, including anInvoiceTotalsItsLines test.
- `Ledger.kt` (291 B) — Kotlin Ledger class with sum() method, LEDGER_NAME constant, and newLedger() factory function. · declares `Ledger`, `fun`
- `LedgerTests.kt` (275 B) — Kotlin tests for Ledger, including aLedgerSumsItsEntries test.
- `Payment.cs` (269 B) — C# Payment class with RetryLimit constant and Settle() method to process an amount. · declares `Payment`
- `Refund.swift` (347 B) — Swift Refund struct with isAllowed() method, refundWindowDays constant, and makeRefund() factory function. · declares `Refund`, `isAllowed`, `makeRefund`

## Structure

- InvoiceTest.java exercises Invoice.java's total() method
- LedgerTests.kt exercises Ledger.kt's sum() method via Ledger or newLedger()
- Each language file (Invoice, Ledger, Payment, Refund) stands independently with no cross-file calls shown

## Rules

- Invoice.MAX_LINES caps invoices at 200 lines
- Payment.RetryLimit caps payment retries at 5
- refundWindowDays fixes the refund window at 30 days

## Where to look

- how invoice totals are computed → `Invoice.java` `total`
- invoice line limit → `Invoice.java` `MAX_LINES`
- ledger summation logic → `Ledger.kt` `sum`
- creating an empty ledger → `Ledger.kt` `newLedger`
- payment retry limit → `Payment.cs` `RetryLimit`
- settling a payment → `Payment.cs` `Settle`
- refund eligibility check → `Refund.swift` `isAllowed`
- refund window length → `Refund.swift` `refundWindowDays`
- invoice test coverage → `InvoiceTest.java` `anInvoiceTotalsItsLines`
- ledger test coverage → `LedgerTests.kt` `aLedgerSumsItsEntries`
