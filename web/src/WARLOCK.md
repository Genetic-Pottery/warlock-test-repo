<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# src

Root source directory holding a small cart API stub and a legacy retry helper, alongside a components subdirectory with Cart UI logic and its tests.

## Files

- `api.ts` (130 B) — Defines BASE_URL constant and fetchCart(id) stub for retrieving a cart by id. · declares `BASE_URL`, `fetchCart`
- `legacy.js` (145 B) — Legacy retry helper: RETRY_LIMIT constant (3) and retry(fn) that calls fn up to RETRY_LIMIT times. · declares `RETRY_LIMIT`, `retry`

## Directories

- `components/` — UI-adjacent Cart class (add, lines, CartLine, CartState, CART_LIMIT) and its tests; go here for cart line/state logic or Cart tests.

## Structure

- Defines BASE_URL constant and fetchCart(id) stub for retrieving a cart by id.
- Legacy retry helper: RETRY_LIMIT constant (3) and retry(fn) that calls fn up to RETRY_LIMIT times.
