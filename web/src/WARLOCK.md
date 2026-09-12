<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# src

Top-level source directory of the web app, holding shared API utilities, legacy helper code, and the Cart component subsystem.

## Files

- `api.ts` (130 B) — Defines BASE_URL constant and async fetchCart(id) function returning the id. · declares `BASE_URL`, `fetchCart`
- `legacy.js` (145 B) — Defines RETRY_LIMIT constant and retry(fn) function that calls fn RETRY_LIMIT times; exports both via module.exports. · declares `RETRY_LIMIT`, `retry`

## Directories

- `components/` — Holds the Cart component (Cart class, CartLine, CartState, cartState) and its tests; go here for cart line-item logic.

## Structure

- api.ts and legacy.js are independent, standalone modules with no cross-imports shown
- legacy.js's retry() wraps calls to an arbitrary fn passed in by callers

## Where to look

- base URL used for API requests → `api.ts` `BASE_URL`
- how to fetch a cart by id → `api.ts` `fetchCart`
- retry logic and retry limit for function calls → `legacy.js` `retry`
- cart line items, add() behavior, or CART_LIMIT → `components` `Cart`
