<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# api

Entry point for the API service, defining the Server struct and Boot() to construct a server bound to a fixed listen address.

## Files

- `server.go` (140 B) — Defines Server struct{addr}, ListenAddr constant "0.0.0.0:8080", and Boot() constructing a Server bound to that address. · declares `Server`, `Boot`

## Directories

- `handlers/` — Route registration and HTTP handlers; go here for how endpoints like /entries/reverse are wired.

## Structure

- Defines Server struct{addr}, ListenAddr constant "0.0.0.0:8080", and Boot() constructing a Server bound to that address.
