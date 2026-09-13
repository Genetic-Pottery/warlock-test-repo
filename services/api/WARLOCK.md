<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# api

The api directory is the entry point for the api service, defining the Server type that boots the service and listens on a fixed address.

## Files

- `server.go` (140 B) — Defines Server struct with addr field, ListenAddr constant ("0.0.0.0:8080"), and Boot() constructor returning a Server initialized with ListenAddr. · declares `Server`, `Boot`

## Directories

- `handlers/` — HTTP route handling with Router type, DefaultPrefix, NewRouter, and Handle; go there for routing or path-validity questions.

## Structure

- Boot constructs a Server initialized with ListenAddr.

## Rules

- ListenAddr is fixed to "0.0.0.0:8080".

## Where to look

- what address the api server listens on → `server.go` `ListenAddr`
- how the api server is started → `server.go` `Boot`
- how routes are matched or versioned → `handlers` `Router`
