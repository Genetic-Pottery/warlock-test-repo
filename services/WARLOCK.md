<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# services

The services directory holds independent service implementations, currently the HTTP api service and the billing domain models, each with its own language and structure.

## Directories

- `api/` — HTTP API service with a Server and ListenAddr; consult for how the service boots or listens for requests.
- `billing/` — Billing domain models (Invoice, Ledger, Payment, Refund) and their tests; consult for invoice, ledger, payment, or refund logic.
