<!-- warlock -->
> Written by a model pass over this directory alone, to be read before its source and to say which source to read. A map, not a specification: check anything you are about to rely on against the files themselves, and where this document and the code disagree, the code is right.

# pipeline

Contains Pipeline.Stage, a single pipeline stage that sums items via run/1, plus its ExUnit tests.

## Files

- `stage.ex` (193 B) — Pipeline.Stage: one pipeline stage; run/1 sums items via Enum.reduce, private normalise/1 passes items through unchanged. · declares `Pipeline.Stage`, `run`, `normalise`
- `stage_test.exs` (272 B) — ExUnit tests for Pipeline.Stage's run/1, covering summing of given items.

## Structure

- Pipeline.Stage: one pipeline stage; run/1 sums items via Enum.reduce, private normalise/1 passes items through unchanged.
- ExUnit tests for Pipeline.Stage's run/1, covering summing of given items.
