<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# services

Collects the backend service components: an API entry point that boots a server and routes requests, and billing types across several languages modeling invoices, ledgers, payments, and refunds.

## Directories

- `api/` — HTTP entry point: Server struct and Boot() bind a listen address; go there for boot or routing questions.
- `billing/` — Multi-language billing domain types (invoices, ledgers, payments, refunds); go there for billing logic or validation questions.
