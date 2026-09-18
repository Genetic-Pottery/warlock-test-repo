<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# legacy

Legacy codec and decoder module providing frame encoding and ID-based decoding, superseded but retained for compatibility.

## Files

- `codec.c` (397 B) — Legacy codec: CODEC_VERSION=4, Frame struct, encode() summing 0..7 plus frame->len and bumping static frames_seen.
- `codec.h` (110 B) — Header declaring CODEC_VERSION constant and encode(struct Frame*) function prototype for the legacy codec interface.
- `decoder.cpp` (263 B) — Legacy Decoder class with decode(id): loops kMaxFrames counting total, returns total > id.

## Structure

- Legacy codec: CODEC_VERSION=4, Frame struct, encode() summing 0..7 plus frame->len and bumping static frames_seen.
- Header declaring CODEC_VERSION constant and encode(struct Frame*) function prototype for the legacy codec interface.
- Legacy Decoder class with decode(id): loops kMaxFrames counting total, returns total > id.
