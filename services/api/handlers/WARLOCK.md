<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# handlers

Defines HTTP routing for the API under the "/v1" prefix, registering routes such as POST /entries/reverse and matching incoming requests to them.

## Files

- `routes.go` (998 B) — Router with DefaultPrefix "/v1": NewRouter registers routes (POST /entries/reverse); Handle(method, path) checks prefix and matches route. · declares `route`, `Router`, `NewRouter`, `Handle`
- `routes_test.go` (871 B) — Tests for route matching: verifies Handle matches method+path, rejects wrong method, missing prefix, and unregistered paths.

## Structure

- Router with DefaultPrefix "/v1": NewRouter registers routes (POST /entries/reverse); Handle(method, path) checks prefix and matches route.
- Tests for route matching: verifies Handle matches method+path, rejects wrong method, missing prefix, and unregistered paths.
