<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# legacy

Legacy codec subsystem providing frame encoding via a C codec and a C++ Decoder class for decoding frames by id.

## Files

- `codec.c` (397 B) — Legacy codec implementation: defines CODEC_VERSION, Frame struct, and encode(), which tallies a fixed loop and increments static frames_seen.
- `codec.h` (110 B) — Header declaring CODEC_VERSION and the encode(Frame*) function signature; guarded by CODEC_H include guard.
- `decoder.cpp` (263 B) — Defines legacy::Decoder class with decode(int id) method and kMaxFrames constant, includes codec.h.

## Structure

- decoder.cpp includes codec.h to access encode() and CODEC_VERSION.
- codec.c implements the encode() function declared in codec.h.
- Decoder::decode() in decoder.cpp is independent logic, not calling encode() directly.

## Rules

- codec.h uses CODEC_H include guard to prevent double inclusion.
- encode() increments frames_seen as a static counter each call.

## Where to look

- how frames are encoded → `codec.c` `encode`
- codec version number → `codec.h` `CODEC_VERSION`
- frame decoding logic → `decoder.cpp` `Decoder`
- Frame struct definition → `codec.c` `Frame`
- max frames constant → `decoder.cpp` `kMaxFrames`
