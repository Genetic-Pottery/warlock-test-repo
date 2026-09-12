<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# handlers

The handlers package defines HTTP routing for the API service, providing the Router type that registers and matches request paths under a versioned prefix.

## Files

- `routes.go` (298 B) — Defines Router struct with prefix field, DefaultPrefix constant ("/v1"), NewRouter() constructor, and Handle(path) method returning whether path is non-empty. · declares `Router`, `NewRouter`, `r`
- `routes_test.go` (312 B) — Test file with TestRouterHandlesNonEmptyPaths verifying Router.Handle behavior on non-empty paths.

## Structure

- NewRouter() constructs a Router with prefix set to DefaultPrefix
- Router.Handle(path) is called on instances returned by NewRouter()
- routes_test.go depends on routes.go's Router and Handle

## Rules

- DefaultPrefix is fixed to "/v1"
- Handle(path) returns true only when path length is greater than 0

## Where to look

- how are HTTP routes registered → `routes.go` `Router`
- what is the default API version prefix → `routes.go` `DefaultPrefix`
- how to create a new router → `routes.go` `NewRouter`
- tests for path handling → `routes_test.go` `TestRouterHandlesNonEmptyPaths`
