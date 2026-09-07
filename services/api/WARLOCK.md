<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# api

The api directory defines the API service's server entry point, providing a Server type and boot function used to start listening for requests.

## Files

- `server.go` (140 B) — Defines Server struct with addr field, ListenAddr constant ("0.0.0.0:8080"), and Boot() constructor returning a new Server. · declares `Server`, `Boot`

## Directories

- `handlers/` — HTTP routing package with Router type and path-matching logic; look here for how request paths are handled or prefixed.

## Structure

- Boot constructs a Server using the ListenAddr constant

## Rules

- ListenAddr is fixed to "0.0.0.0:8080"

## Where to look

- what address does the server listen on → `server.go` `ListenAddr`
- how to start the server → `server.go` `Boot`
- how are HTTP routes handled → `handlers` `Router`
