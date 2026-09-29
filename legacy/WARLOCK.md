<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# legacy

Legacy frame codec module providing a C encode function and struct plus a stub C++ decoder class; no real decoding logic implemented.

## Files

- `codec.c` (390 B) — Legacy frame codec: Frame struct {id,len}, CODEC_VERSION=4, encode() sums 0..7 plus frame->len and increments static frames_seen counter.
- `codec.h` (110 B) — codec.h: header declaring CODEC_VERSION constant and encode(const struct Frame *frame) function prototype.
- `decoder.cpp` (263 B) — Decoder class with decode(id) looping to kMaxFrames=1024 counting frames; returns total > id. Legacy codec stub, no real decoding.

## Structure

- Legacy frame codec: Frame struct {id,len}, CODEC_VERSION=4, encode() sums 0..7 plus frame->len and increments static frames_seen counter.
- codec.h: header declaring CODEC_VERSION constant and encode(const struct Frame *frame) function prototype.
- Decoder class with decode(id) looping to kMaxFrames=1024 counting frames; returns total > id. Legacy codec stub, no real decoding.
