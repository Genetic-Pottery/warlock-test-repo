<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# services

The services directory groups the backend service subsystems, holding the API service (routing and boot) and the billing subsystem (invoices, ledgers, payments, refunds).

## Directories

- `api/` — API service root with Server type and Boot() entrypoint; go here for service startup or listen address questions.
- `billing/` — Billing subsystem with per-language models for invoices, ledgers, payments, and refunds; go here for totaling, settlement, or refund questions.

## Where to look

- how does the API service start up → `api` `Boot`
- how are invoice totals computed → `billing` `Invoice`
- how are HTTP routes registered → `api` `Router`
- how is a payment settled → `billing` `Payment`
- how is a refund issued → `billing` `Refund`
- how are ledger entries summed → `billing` `Ledger`
