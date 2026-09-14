<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# handlers

Provides HTTP routing via a Router struct that matches request paths against a configurable prefix, defaulting to "/v1".

## Files

- `routes.go` (298 B) — Defines Router struct with prefix field, DefaultPrefix ("/v1") constant, NewRouter constructor, and Handle(path) route matcher. · declares `Router`, `NewRouter`, `Handle`
- `routes_test.go` (312 B) — Test file for router behavior; contains TestRouterHandlesNonEmptyPaths verifying routing on non-empty paths.

## Structure

- Defines Router struct with prefix field, DefaultPrefix ("/v1") constant, NewRouter constructor, and Handle(path) route matcher.
- Test file for router behavior; contains TestRouterHandlesNonEmptyPaths verifying routing on non-empty paths.
