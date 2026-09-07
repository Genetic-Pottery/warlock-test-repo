<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# services

The services directory groups independent backend services, each with its own domain: the api service handles HTTP request routing, and the billing service holds core billing domain types for invoices, ledgers, payments, and refunds.

## Directories

- `api/` — API service's server entry point, with Server type, ListenAddr, Boot function, and request handlers; look here for how the server starts and routes requests.
- `billing/` — Billing domain types for invoices, ledgers, payments, and refunds across multiple languages; look here for billing totals, limits, retries, or refund rules.

## Where to look

- how does the API server start and what address does it bind → `api` `Boot`
- how is an invoice total computed or capped → `billing` `Invoice`
- how are payments retried or settled → `billing` `Payment`
- what is the refund eligibility window → `billing` `Refund`
