<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# components

Holds the Cart component logic for the shopping cart, defining the Cart class, its line-item model, and state helpers used by the UI.

## Files

- `Cart.test.tsx` (370 B) — Unit tests for Cart: verifies add() rejects duplicate SKUs and accepts new lines.
- `Cart.tsx` (459 B) — Defines CART_LIMIT constant, CartLine interface, CartState type, the Cart class (add method), and cartState() function. · declares `CART_LIMIT`, `CartLine`, `CartState`, `Cart`, `cartState`

## Structure

- Cart.test.tsx imports and exercises Cart from Cart.tsx
- cartState() takes a Cart instance and returns a CartState

## Rules

- CART_LIMIT is set to 50
- Cart.add() returns false if the sku already exists, true otherwise

## Where to look

- how duplicate items are prevented in the cart → `Cart.tsx` `add`
- maximum number of items allowed in a cart → `Cart.tsx` `CART_LIMIT`
- tests for cart add behavior → `Cart.test.tsx` `Cart`
