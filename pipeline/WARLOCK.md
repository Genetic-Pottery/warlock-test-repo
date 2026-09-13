<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# pipeline

The pipeline directory holds the Pipeline subsystem's stage logic, defining a single processing stage that reduces a list of items into a summed value.

## Files

- `stage.ex` (193 B) — Defines Pipeline.Stage with run/1, which reduces items by summing them, and a private normalise/1 helper that currently passes items through unchanged. · declares `Pipeline.Stage`, `run`, `normalise`
- `stage_test.exs` (272 B) — Pipeline.StageTest, tests for run/1 covering that it sums the items it is given.

## Structure

- run/1 uses Enum.reduce to sum items, starting from an accumulator of 0
- stage_test.exs exercises Pipeline.Stage's run/1 function

## Where to look

- how items are combined in a pipeline stage → `stage.ex` `run`
