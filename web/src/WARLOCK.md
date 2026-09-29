<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# src

Top-level source directory holding a stubbed cart-fetching API client and a legacy retry utility, alongside the components subdirectory implementing the Cart itself.

## Files

- `api.ts` (130 B) — Defines BASE_URL constant and async fetchCart(id) stub returning the given id; not yet implemented. · declares `BASE_URL`, `fetchCart`
- `legacy.js` (145 B) — Defines RETRY_LIMIT (3) and retry(fn), a loop calling fn up to RETRY_LIMIT times; exports both via module.exports. · declares `RETRY_LIMIT`, `retry`

## Directories

- `components/` — Holds the Cart component and its tests for cart state management, including adding unique-sku line items.

## Structure

- Defines BASE_URL constant and async fetchCart(id) stub returning the given id; not yet implemented.
- Defines RETRY_LIMIT (3) and retry(fn), a loop calling fn up to RETRY_LIMIT times; exports both via module.exports.
