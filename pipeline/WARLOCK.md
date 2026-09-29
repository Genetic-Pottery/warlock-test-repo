<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# pipeline

Contains Pipeline.Stage, a module that runs a processing stage by summing the items it receives, along with its test suite.

## Files

- `stage.ex` (193 B) — Defines Pipeline.Stage with run/1, summing items via Enum.reduce, and an unused private normalise/1 stub that returns its input unchanged. · declares `Pipeline.Stage`, `run`, `normalise`
- `stage_test.exs` (272 B) — Test suite for Pipeline.Stage's run/1, verifying it sums the items it receives.

## Structure

- Defines Pipeline.Stage with run/1, summing items via Enum.reduce, and an unused private normalise/1 stub that returns its input unchanged.
- Test suite for Pipeline.Stage's run/1, verifying it sums the items it receives.
