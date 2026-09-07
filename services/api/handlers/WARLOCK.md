<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# handlers

Package handlers defines HTTP routing for the API service, providing a Router type that registers and matches request paths under a versioned prefix.

## Files

- `routes.go` (298 B) — Defines Router struct with prefix field, DefaultPrefix constant ("/v1"), NewRouter() constructor, and Handle(path string) bool method checking non-empty paths. · declares `Router`, `NewRouter`, `r`
- `routes_test.go` (312 B) — Tests Router.Handle, including TestRouterHandlesNonEmptyPaths verifying non-empty paths are handled.

## Structure

- routes_test.go calls NewRouter and Router.Handle defined in routes.go
- NewRouter initializes Router with DefaultPrefix

## Where to look

- what path prefix does the API use → `routes.go` `DefaultPrefix`
- how are routes matched or handled → `routes.go` `Handle`
- test coverage for routing → `routes_test.go` `TestRouterHandlesNonEmptyPaths`
