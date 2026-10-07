# warlock-test-repo

A fixture repository for sanity-testing warlock. Nothing here is real code.

The directories are shaped to exercise different parts of a pass: several
languages, one very large file, one binary, one deep nesting, and one directory
big enough that a request has to fit rather than simply arrive.

## Test suites

Four suites sit at the repository root. `./run.sh` runs the two that cost
nothing; `./run.sh --paid` runs all four with `check.sh` first, because
`check.sh` is what pacts the fixture that the other three read.

| Suite | What it asserts | What it spends | Prerequisites |
| --- | --- | --- | --- |
| `surfaces.sh` | That `stale` and `fresh` are the two halves of one ledger, in prose and in `--json`; the exit-status table a script reads; and the `brief`/`push` front door, where every `push` carries `--dry-run`. | Nothing. Seconds and no tokens: neither `pact` nor `refresh` appears in a direction that walks a directory. | A pacted fixture, which only a paid `check.sh` provides. It exits 2 without `.warlock/pacts.toml`, and it also needs the `warlock` binary and `jq`. |
| `scopes.sh` | The scope and sigil boundary: that a mutating key is refused across a closed scope with nothing spent, that the sigil is what opens it, that the nearest scope wins, and that clearing the sigils closes every scope again. Assertions are exit statuses and `--json` fields, never sentences. | Nothing. `pact` and `refresh` appear only in the refusal direction, where the gate answers before a directory is walked. | A pacted fixture, which only a paid `check.sh` provides: it wants `.warlock/pacts.toml` with `legacy`, `services`, `services/api` and `services/api/handlers` named in it, plus the `warlock` binary and `jq`. |
| `check.sh` | That a full pact reaches every document: the symbols the language table must surface, a binary file named rather than described, and a document present for each directory. Assertions are greps, never exact text. | One model pass per directory, over the 16 directories it pacts. | The `warlock` binary. It needs no pacted fixture, because it is what pacts one; it re-pacts in place, so it leaves the tracked documents modified. `./check.sh --no-pact` checks the documents already on disk and spends nothing. |
| `incremental.sh` | What a refresh reuses: that a changed, new or deleted file costs its own line and no other, byte for byte, and that editing a `.warlockignore` stales nothing. Every refresh is aimed at `engine/core`. | A settling refresh, then one model pass per scenario — each a file pass plus that directory's synthesis. | A fixture that is already pacted by `check.sh`, an existing `engine/core/.warlock.md`, and a clean tree, because it restores source files with git. |

No suite writes to Linear. `surfaces.sh` is the only one that reaches the brief
commands, and every `push` there carries `--dry-run`, so no socket is opened and
nothing arrives at a board.

The round trip of `brief` → `push` → `draft` → `pull` is exercised by hand
against the GEN board. `docs/` holds the briefs those runs left behind.
