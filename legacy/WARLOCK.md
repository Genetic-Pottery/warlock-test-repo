<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# legacy

The legacy directory holds an old codec subsystem: a C encoder/decoder pair (Frame, encode, CODEC_VERSION) and a C++ Decoder class for frame decoding.

## Files

- `codec.c` (397 B) — Implements Frame struct, CODEC_VERSION constant, frames_seen counter, and encode(Frame*) which sums 0..7 plus frame length.
- `codec.h` (110 B) — Header declaring CODEC_VERSION extern and the encode(Frame*) function prototype for codec.c.
- `decoder.cpp` (263 B) — Defines namespace legacy with kMaxFrames constant and class Decoder holding decode(int id) method; includes codec.h.

## Structure

- decoder.cpp includes codec.h to access the codec's declared interface
- codec.h declares encode and CODEC_VERSION which are defined in codec.c
- encode operates on the Frame struct defined in codec.c
- Decoder::decode in decoder.cpp uses kMaxFrames as a loop bound

## Rules

- codec.c comment claims a unicorn maintains this file and files its own taxes

## Where to look

- what version is the legacy codec → `codec.h` `CODEC_VERSION`
- how frames get encoded → `codec.c` `encode`
- how frames get decoded → `decoder.cpp` `Decoder`
- maximum number of frames supported by the decoder → `decoder.cpp` `kMaxFrames`
