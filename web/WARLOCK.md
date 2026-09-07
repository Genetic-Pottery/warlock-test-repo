<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# web

Web app root directory containing the src subdirectory, which holds the cart API helper, legacy retry utility, and cart component logic for the shopping cart UI.

## Directories

- `src/` — Top-level source: fetchCart API helper, legacy retry utility, and components subdirectory with Cart class and cart limit logic.

## Where to look

- fetching a cart by id → `src` `fetchCart`
- retrying a function call a fixed number of times → `src` `retry`
- cart add/duplicate item logic → `src` `Cart`
- maximum items allowed in a cart → `src` `CART_LIMIT`
