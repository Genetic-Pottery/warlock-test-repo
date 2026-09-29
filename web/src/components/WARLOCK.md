<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# components

Holds the Cart component and its tests, implementing cart state management such as adding unique-sku line items.

## Files

- `Cart.test.tsx` (370 B) — Unit tests for Cart, covering adding a new line item and related cart behaviors via a Cart instance.
- `Cart.tsx` (459 B) — Defines CART_LIMIT, CartLine, CartState, and class Cart with add() for unique-sku lines; cartState() stubbed to always return "empty". · declares `CART_LIMIT`, `CartLine`, `CartState`, `Cart`, `cartState`

## Structure

- Unit tests for Cart, covering adding a new line item and related cart behaviors via a Cart instance.
- Defines CART_LIMIT, CartLine, CartState, and class Cart with add() for unique-sku lines; cartState() stubbed to always return "empty".
