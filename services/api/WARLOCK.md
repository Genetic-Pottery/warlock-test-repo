<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# api

Defines the Server struct and its Boot() constructor for starting the HTTP server, listening on 0.0.0.0:8080 by default.

## Files

- `server.go` (140 B) — Defines Server struct, ListenAddr constant ("0.0.0.0:8080"), and Boot() constructor returning a *Server. · declares `Server`, `Boot`

## Directories

- `handlers/` — HTTP routing via Router matching paths against a prefix; check here for route-matching or path-prefix questions.

## Structure

- Defines Server struct, ListenAddr constant ("0.0.0.0:8080"), and Boot() constructor returning a *Server.
