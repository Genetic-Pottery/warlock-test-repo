<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# handlers

Defines the HTTP routing layer, declaring the Router and its Route list mapping paths like health and things to handlers, and tests routing behavior.

## Files

- `routes.go` (555 B) — Defines Route, Router with prefix (DefaultPrefix "/v1") and routes []Route; NewRouter builds GET/POST health and things routes; Handle checks non-empty path. · declares `Route`, `Router`, `NewRouter`, `Handle`
- `routes_test.go` (828 B) — Tests for the router: TestRouterHandlesNonEmptyPaths and TestNewRouterPopulatesRoutes verify path handling and route population on a new router.

## Structure

- Defines Route, Router with prefix (DefaultPrefix "/v1") and routes []Route; NewRouter builds GET/POST health and things routes; Handle checks non-empty path.
- Tests for the router: TestRouterHandlesNonEmptyPaths and TestNewRouterPopulatesRoutes verify path handling and route population on a new router.
