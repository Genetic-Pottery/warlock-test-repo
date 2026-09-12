<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# billing

Billing subsystem covering invoices, ledgers, payments, and refunds, with per-language models for totaling invoice lines, summing ledger entries, settling payments, and issuing refunds.

## Files

- `Invoice.java` (326 B) — Invoice class with MAX_LINES constant and total() method summing invoice lines. · declares `Invoice`
- `InvoiceTest.java` (346 B) — InvoiceTest with anInvoiceTotalsItsLines test; test bodies elided.
- `Ledger.kt` (291 B) — Ledger class with sum() method, LEDGER_NAME constant, and newLedger() factory function. · declares `Ledger`, `fun`
- `LedgerTests.kt` (275 B) — LedgerTests with aLedgerSumsItsEntries test; test bodies elided.
- `Payment.cs` (269 B) — Payment class with RetryLimit constant and Settle() method for processing amounts. · declares `Payment`
- `Refund.swift` (347 B) — Refund struct with isAllowed() method, refundWindowDays constant, and makeRefund() factory function. · declares `Refund`, `isAllowed`, `makeRefund`

## Structure

- Ledger.kt's newLedger() constructs a Ledger with empty entries.
- Refund.swift's makeRefund() constructs a Refund from an amount.
- InvoiceTest.java tests Invoice.java's total().
- LedgerTests.kt tests Ledger.kt's sum().

## Rules

- Invoice.MAX_LINES caps invoices at 200 lines.
- Payment.RetryLimit caps retries at 5.
- refundWindowDays fixes the refund window at 30 days.

## Where to look

- how invoice totals are computed → `Invoice.java` `total`
- invoice line limit → `Invoice.java` `MAX_LINES`
- how ledger entries are summed → `Ledger.kt` `sum`
- creating a new empty ledger → `Ledger.kt` `newLedger`
- payment retry limit → `Payment.cs` `RetryLimit`
- settling a payment → `Payment.cs` `Settle`
- refund window length → `Refund.swift` `refundWindowDays`
- creating a refund → `Refund.swift` `makeRefund`
- invoice test coverage → `InvoiceTest.java` `anInvoiceTotalsItsLines`
- ledger test coverage → `LedgerTests.kt` `aLedgerSumsItsEntries`
