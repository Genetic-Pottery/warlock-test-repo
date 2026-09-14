<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# services

Holds service-layer subsystems, grouping the HTTP API server and its routing from the billing domain types for invoices, ledgers, payments, and refunds.

## Directories

- `api/` — HTTP server startup (Server, Boot) and route-matching handlers; check here for server config or path-prefix routing questions.
- `billing/` — Billing domain types and tests for Invoice, Ledger, Payment, and Refund across Java, Kotlin, C#, and Swift; check here for billing logic questions.
