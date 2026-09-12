<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# web

The web app directory, containing shared API utilities, legacy helper code, and the src subdirectory holding the app's source including the Cart component subsystem.

## Directories

- `src/` — Holds shared API utilities (BASE_URL, fetchCart), legacy retry() helper, and the Cart component subsystem; go here for cart line-item logic or API/retry helpers.

## Where to look

- base URL used for API requests → `src` `BASE_URL`
- how to fetch a cart by id → `src` `fetchCart`
- retry logic and retry limit for function calls → `src` `retry`
- cart line items, add() behavior, or CART_LIMIT → `src` `Cart`
