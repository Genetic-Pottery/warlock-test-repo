<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# pipeline

This is the pipeline directory, holding Pipeline.Stage, a single stage of a processing pipeline that sums a list of items and its accompanying test suite.

## Files

- `stage.ex` (193 B) — Defines Pipeline.Stage with run/1, which reduces a list of items into a sum via Enum.reduce, and a private normalise/1 that returns its item unchanged. · declares `Pipeline.Stage`, `run`, `normalise`
- `stage_test.exs` (272 B) — Defines Pipeline.StageTest, testing run/1, including a case that it sums the items it is given.

## Structure

- Pipeline.StageTest exercises Pipeline.Stage.run/1
- run/1 calls Enum.reduce to accumulate item sums
- normalise/1 is defined but not called within run/1

## Where to look

- how items are summed in the pipeline → `stage.ex` `run`
- tests for the pipeline stage's summing behavior → `stage_test.exs` `run/1`
