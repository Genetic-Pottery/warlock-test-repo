<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# handlers

Provides HTTP routing for the API, defining a Router that checks request paths against a versioned prefix before dispatching.

## Files

- `routes.go` (298 B) — Defines Router type with prefix field, DefaultPrefix ("/v1") constant, NewRouter constructor, and Handle(path) path-check method. · declares `Router`, `NewRouter`, `Handle`
- `routes_test.go` (312 B) — Test file with TestRouterHandlesNonEmptyPaths, verifying the router correctly handles non-empty request paths.

## Structure

- Defines Router type with prefix field, DefaultPrefix ("/v1") constant, NewRouter constructor, and Handle(path) path-check method.
- Test file with TestRouterHandlesNonEmptyPaths, verifying the router correctly handles non-empty request paths.
