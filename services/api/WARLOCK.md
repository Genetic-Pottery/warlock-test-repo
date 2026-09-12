<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# api

The api directory is the root of the API service, providing the Server type and boot entrypoint for the HTTP service, with routing delegated to the handlers subpackage.

## Files

- `server.go` (140 B) — Defines Server struct with addr field, ListenAddr constant ("0.0.0.0:8080"), and Boot() constructor returning a Server bound to ListenAddr. · declares `Server`, `Boot`

## Directories

- `handlers/` — HTTP routing via Router type and path matching under a versioned prefix; go here for route registration or path-handling questions.

## Structure

- Boot() constructs a Server with addr set to ListenAddr

## Rules

- ListenAddr is fixed to "0.0.0.0:8080"

## Where to look

- how does the service start up → `server.go` `Boot`
- what address does the server listen on → `server.go` `ListenAddr`
- how are HTTP routes registered → `handlers` `Router`
