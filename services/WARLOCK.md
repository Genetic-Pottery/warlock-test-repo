<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# services

Parent directory grouping backend service code, holding the top-level api service and the multi-language billing domain types.

## Directories

- `api/` — Server struct, ListenAddr, and Boot() entry point; go there for how the service starts and binds its address.
- `billing/` — Invoice, Ledger, Payment, and Refund types with their summing/validation logic; go there for billing domain questions.
