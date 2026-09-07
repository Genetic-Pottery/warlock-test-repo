<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# legacy

Legacy codec module providing frame encoding via C code and frame decoding via a C++ Decoder class, both keyed to CODEC_VERSION 4.

## Files

- `codec.c` (397 B) — Defines CODEC_VERSION (4), static frames_seen counter, Frame struct, and encode() which sums 0..7 plus frame->len and increments frames_seen.
- `codec.h` (110 B) — Header declaring extern CODEC_VERSION and the encode(const struct Frame *frame) prototype.
- `decoder.cpp` (263 B) — Defines namespace legacy with kMaxFrames (1024) and class Decoder holding decode(int id), which returns true if id is less than kMaxFrames.

## Structure

- decoder.cpp includes codec.h to access CODEC_VERSION and encode declarations
- codec.c implements the encode() function declared in codec.h
- codec.h is the shared interface between codec.c and decoder.cpp

## Rules

- codec.h declares CODEC_VERSION as extern const int, defined once in codec.c
- Decoder::decode loops up to kMaxFrames (1024) before comparing to id

## Where to look

- how frames are encoded → `codec.c` `encode`
- current codec version number → `codec.h` `CODEC_VERSION`
- how frames are decoded → `decoder.cpp` `Decoder`
- max frame limit for decoding → `decoder.cpp` `kMaxFrames`
