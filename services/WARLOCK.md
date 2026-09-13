<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# services

The services directory groups independent service subsystems: the api entry point that boots the HTTP server and its handlers, and the billing domain covering invoices, ledgers, payments, and refunds.

## Directories

- `api/` — Api service entry point with the Server type and Boot(); go there for how the service starts, its listen address, or routing questions.
- `billing/` — Core billing domain types (Invoice, Ledger, Payment, Refund) with totaling, retry, and refund-window logic and tests; go there for billing rules or calculations.

## Where to look

- how the api service is started or what address it listens on → `api` `Boot`
- how billing invoices, ledgers, payments, or refunds are limited or computed → `billing` `Invoice`
