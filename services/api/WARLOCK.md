<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# api

Top-level API service defining the Server struct and Boot() entry point that binds a server to a fixed address, with routing delegated to handlers.

## Files

- `server.go` (140 B) — Defines Server struct{addr}, ListenAddr constant "0.0.0.0:8080", and Boot() constructing a Server bound to that address. · declares `Server`, `Boot`

## Directories

- `handlers/` — HTTP routing: Router, Route list, and path handlers for endpoints like health and things — go there for routing/dispatch questions.

## Structure

- Defines Server struct{addr}, ListenAddr constant "0.0.0.0:8080", and Boot() constructing a Server bound to that address.
