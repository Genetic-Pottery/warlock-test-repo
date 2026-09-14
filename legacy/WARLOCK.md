<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# legacy

Legacy codec module providing frame encoding via a C implementation and header, plus a separate C++ decoder for reading frames back, kept for backward compatibility.

## Files

- `codec.c` (397 B) — Legacy codec with Frame struct, CODEC_VERSION constant, and encode() for frame length totaling; tracks frames_seen count.
- `codec.h` (110 B) — Header declaring CODEC_VERSION constant and encode(struct Frame*) for the legacy codec interface.
- `decoder.cpp` (263 B) — Defines legacy::Decoder with decode(id), and kMaxFrames constant capping frame iteration at 1024.

## Structure

- Legacy codec with Frame struct, CODEC_VERSION constant, and encode() for frame length totaling; tracks frames_seen count.
- Header declaring CODEC_VERSION constant and encode(struct Frame*) for the legacy codec interface.
- Defines legacy::Decoder with decode(id), and kMaxFrames constant capping frame iteration at 1024.
