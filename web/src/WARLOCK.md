<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# src

Top-level source directory of the web app, holding a fetchCart API helper, a legacy retry utility, and the components subdirectory implementing the shopping cart UI logic.

## Files

- `api.ts` (130 B) — Exports BASE_URL constant and async fetchCart(id) function that returns the given cart id. · declares `BASE_URL`, `fetchCart`
- `legacy.js` (145 B) — Legacy CommonJS module exporting RETRY_LIMIT constant (3) and retry(fn) function that calls fn up to RETRY_LIMIT times. · declares `RETRY_LIMIT`, `retry`

## Directories

- `components/` — Cart component logic: Cart class, CartLine/CartState types, CART_LIMIT; go here for cart add behavior or item limits.

## Structure

- api.ts and legacy.js are independent, standalone modules with no imports between them
- legacy.js's retry() calls the passed-in fn up to RETRY_LIMIT times in a loop

## Rules

- legacy.js sets RETRY_LIMIT to 3

## Where to look

- fetching a cart by id → `api.ts` `fetchCart`
- base URL used for API requests → `api.ts` `BASE_URL`
- retrying a function call a fixed number of times → `legacy.js` `retry`
- how many times an operation is retried → `legacy.js` `RETRY_LIMIT`
- cart add/duplicate item logic → `components` `Cart`
- maximum items allowed in a cart → `components` `CART_LIMIT`
