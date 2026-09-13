<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# handlers

Defines HTTP route handling for the api service, providing the Router type used to register and match request paths under a versioned prefix.

## Files

- `routes.go` (298 B) — Defines Router struct with prefix field, DefaultPrefix constant ("/v1"), NewRouter() constructor, and Handle(path string) bool method checking non-empty paths. · declares `Router`, `NewRouter`, `r`
- `routes_test.go` (312 B) — Test file with TestRouterHandlesNonEmptyPaths verifying Router.Handle behavior on non-empty paths.

## Structure

- NewRouter constructs a Router initialized with DefaultPrefix.
- Handle is a method on Router that reports whether a given path is non-empty.
- routes_test.go tests the Handle method of Router.

## Rules

- DefaultPrefix is fixed to "/v1".
- Handle returns true only when the path length is greater than 0.

## Where to look

- what prefix does the API version its routes with → `routes.go` `DefaultPrefix`
- how to check if a path is valid for routing → `routes.go` `Handle`
- tests covering router path handling → `routes_test.go` `TestRouterHandlesNonEmptyPaths`
