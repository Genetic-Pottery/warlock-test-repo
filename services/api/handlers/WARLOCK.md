<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# handlers

Provides HTTP routing logic for the API service, defining a Router that matches request paths against a configured prefix.

## Files

- `routes.go` (298 B) — Defines Router struct with prefix field, NewRouter() constructor using DefaultPrefix ("/v1"), and Handle(path string) bool for route matching. · declares `Router`, `NewRouter`, `Handle`
- `routes_test.go` (312 B) — Test file for router; TestRouterHandlesNonEmptyPaths checks handling of non-empty request paths.

## Structure

- Defines Router struct with prefix field, NewRouter() constructor using DefaultPrefix ("/v1"), and Handle(path string) bool for route matching.
- Test file for router; TestRouterHandlesNonEmptyPaths checks handling of non-empty request paths.
