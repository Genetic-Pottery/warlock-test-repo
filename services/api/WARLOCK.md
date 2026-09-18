<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# api

The api directory implements the HTTP API service, providing a Server that boots and listens for requests on a configured address.

## Files

- `server.go` (140 B) — Defines Server struct, ListenAddr constant (0.0.0.0:8080), and Boot() constructor returning a *Server. · declares `Server`, `Boot`

## Directories

- `handlers/` — Holds HTTP routing logic (Router, path prefix matching); consult for how request paths are matched or handled.

## Structure

- Defines Server struct, ListenAddr constant (0.0.0.0:8080), and Boot() constructor returning a *Server.
