<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# src

Top-level source directory holding API access and legacy retry utilities used elsewhere in the codebase.

## Files

- `api.ts` (130 B) — Defines BASE_URL constant and async fetchCart(id) stub returning the id as a string. · declares `BASE_URL`, `fetchCart`
- `legacy.js` (145 B) — Defines RETRY_LIMIT constant and retry(fn) helper that calls fn up to RETRY_LIMIT times; exports { retry, RETRY_LIMIT }. · declares `RETRY_LIMIT`, `retry`

## Directories

- `components/` — Holds the Cart component, its state types (CartLine, CartState), CART_LIMIT constant, and unit tests.

## Structure

- Defines BASE_URL constant and async fetchCart(id) stub returning the id as a string.
- Defines RETRY_LIMIT constant and retry(fn) helper that calls fn up to RETRY_LIMIT times; exports { retry, RETRY_LIMIT }.
