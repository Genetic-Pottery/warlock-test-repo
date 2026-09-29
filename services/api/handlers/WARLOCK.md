<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# handlers

Defines HTTP route registration for the API, including the router type, default path prefix, and endpoint handlers such as reverse entries.

## Files

- `routes.go` (522 B) — Defines route/Router types with DefaultPrefix "/v1" and NewRouter() registering POST /entries/reverse; Handle(path) just checks non-empty path, not actual routing. · declares `route`, `Router`, `NewRouter`, `Handle`
- `routes_test.go` (312 B) — Test file with TestRouterHandlesNonEmptyPaths, verifying router behavior for non-empty request paths.

## Structure

- Defines route/Router types with DefaultPrefix "/v1" and NewRouter() registering POST /entries/reverse; Handle(path) just checks non-empty path, not actual routing.
- Test file with TestRouterHandlesNonEmptyPaths, verifying router behavior for non-empty request paths.
