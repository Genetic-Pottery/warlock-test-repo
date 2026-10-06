<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# services

Top-level directory grouping backend service subsystems, currently the API server and billing domain types.

## Directories

- `api/` — Server binding and boot logic on 0.0.0.0:8080; check here for routing or server startup questions.
- `billing/` — Invoice, Ledger, Payment, and Refund types across Java, Kotlin, C#, Swift; check here for billing domain logic.
