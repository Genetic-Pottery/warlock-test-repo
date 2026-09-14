<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# components

Holds the Cart component: cart state management including CartLine and CartState types, the CART_LIMIT constant, and its unit tests.

## Files

- `Cart.test.tsx` (370 B) — Unit tests for Cart, covering adding a line item not already present and related cart behavior.
- `Cart.tsx` (459 B) — Defines Cart class with add(), CartLine and CartState types, cartState() helper, and CART_LIMIT constant. · declares `CART_LIMIT`, `CartLine`, `CartState`, `Cart`, `cartState`

## Structure

- Defines Cart class with add(), CartLine and CartState types, cartState() helper, and CART_LIMIT constant.
- Unit tests for Cart, covering adding a line item not already present and related cart behavior.
