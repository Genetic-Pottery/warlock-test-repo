<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# src

Top-level source directory of the web app, holding a small API client and a legacy retry utility alongside the components subsystem for the shopping cart UI.

## Files

- `api.ts` (130 B) — Defines BASE_URL constant and async fetchCart(id) function that returns the given id. · declares `BASE_URL`, `fetchCart`
- `legacy.js` (145 B) — Defines RETRY_LIMIT constant and retry(fn) helper that calls fn up to RETRY_LIMIT times; exports both via module.exports. · declares `RETRY_LIMIT`, `retry`

## Directories

- `components/` — Cart component logic: Cart class, CartLine/CartState types, cartState() derivation, and its tests.

## Structure

- retry() in legacy.js loops up to RETRY_LIMIT times invoking the passed fn
- fetchCart in api.ts is independent of legacy.js and components, just returning the given id

## Rules

- retry(fn) calls fn exactly RETRY_LIMIT times, currently 3

## Where to look

- base URL used for API requests → `api.ts` `BASE_URL`
- fetching a cart by id → `api.ts` `fetchCart`
- retrying a failing function call → `legacy.js` `retry`
- how many times an operation is retried → `legacy.js` `RETRY_LIMIT`
- cart item logic and UI state → `components` `Cart`
