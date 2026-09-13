<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# components

Holds the Cart component logic for the shopping cart, providing the Cart class, its line-item types, and cart state derivation used by the app's cart UI.

## Files

- `Cart.test.tsx` (370 B) — Tests for Cart, including that adding a line not already present succeeds.
- `Cart.tsx` (459 B) — Defines CART_LIMIT constant, CartLine and CartState types, the Cart class with add(), and cartState() function. · declares `CART_LIMIT`, `CartLine`, `CartState`, `Cart`, `cartState`

## Structure

- Cart.test.tsx imports and exercises the Cart class defined in Cart.tsx
- cartState() in Cart.tsx takes a Cart instance and returns a CartState

## Rules

- Cart.add() rejects a line whose sku is already present, returning false

## Where to look

- how many items a cart can hold → `Cart.tsx` `CART_LIMIT`
- shape of a cart line item → `Cart.tsx` `CartLine`
- how cart empty/filled status is determined → `Cart.tsx` `cartState`
- tests for adding items to the cart → `Cart.test.tsx` `Cart`
