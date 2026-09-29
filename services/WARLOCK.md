<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# services

Collection of backend service components: the API entry point binding a server to a listen address, and billing domain types across multiple languages modeling invoices, ledgers, payments, and refunds.

## Directories

- `api/` — API service entry point defining Server and Boot(); go here for the listen address or how the server is constructed.
- `billing/` — Multi-language billing types (Invoice, Ledger, Payment, Refund); go here for how amounts are summed or validated.
