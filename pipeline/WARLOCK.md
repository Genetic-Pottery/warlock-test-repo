<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# pipeline

Defines a single stage of the pipeline, responsible for processing a collection of items into a reduced result.

## Files

- `stage.ex` (193 B) — Defines Pipeline.Stage with run/1, which sums items via Enum.reduce, and private normalise/1 helper. · declares `Pipeline.Stage`, `run`, `normalise`
- `stage_test.exs` (272 B) — Tests for Pipeline.Stage.run/1, verifying it sums the items it is given.

## Structure

- stage_test.exs calls Pipeline.Stage.run/1 defined in stage.ex
- run/1 in stage.ex uses Enum.reduce and could call normalise/1

## Where to look

- how items are summed in a pipeline stage → `stage.ex` `run`
- tests verifying stage summation behavior → `stage_test.exs` `run/1`
