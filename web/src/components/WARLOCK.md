<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# components

Holds the Cart component logic: a Cart class managing line items with add-by-sku uniqueness, plus its unit tests.

## Files

- `Cart.test.tsx` (370 B) — Tests for Cart, checking that adding a line not already present succeeds; other cases elided.
- `Cart.tsx` (459 B) — Defines CART_LIMIT constant, CartLine interface, CartState type, Cart class with add(), and cartState() function. · declares `CART_LIMIT`, `CartLine`, `CartState`, `Cart`, `cartState`

## Structure

- Cart.test.tsx imports and exercises Cart from Cart.tsx
- cartState() takes a Cart and returns a CartState

## Where to look

- how duplicate SKUs are handled when adding to cart → `Cart.tsx` `add`
- maximum number of items allowed in cart → `Cart.tsx` `CART_LIMIT`
- test coverage for adding cart lines → `Cart.test.tsx` `Cart`
