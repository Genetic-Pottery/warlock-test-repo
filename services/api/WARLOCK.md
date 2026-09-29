<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# api

Entry point for the API service, defining the Server type and boot logic that binds it to a fixed listen address.

## Files

- `server.go` (140 B) — Defines Server struct{addr}, ListenAddr constant "0.0.0.0:8080", and Boot() constructing a Server bound to that address. · declares `Server`, `Boot`

## Directories

- `handlers/` — HTTP routing: Router checks paths against a versioned prefix before dispatch; go there for routing or path-handling questions.

## Structure

- Defines Server struct{addr}, ListenAddr constant "0.0.0.0:8080", and Boot() constructing a Server bound to that address.
