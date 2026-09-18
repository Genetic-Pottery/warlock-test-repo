<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# components

Holds UI-adjacent Cart logic: a Cart class managing cart lines and state, plus its tests.

## Files

- `Cart.test.tsx` (370 B) — Test suite for Cart: covers adding a not-yet-present line item to a new Cart instance.
- `Cart.tsx` (459 B) — Cart.tsx: Cart class (add, lines) with CartLine, CartState, CART_LIMIT=50 constant, and cartState() helper. · declares `CART_LIMIT`, `CartLine`, `CartState`, `Cart`, `cartState`

## Structure

- Cart.tsx: Cart class (add, lines) with CartLine, CartState, CART_LIMIT=50 constant, and cartState() helper.
- Test suite for Cart: covers adding a not-yet-present line item to a new Cart instance.
