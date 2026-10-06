# Cover warlock 0.1.1's untested surfaces and say what each suite spends

`check.sh`, `scopes.sh` and `incremental.sh` exercise `pact`, `refresh`, `scope`, `check`, `config` and `unpact`. Nothing exercises the rest. `warlock stale` and `warlock fresh` are never run — `incremental.sh` calls `stale` only as an emptiness test at line 105, and no suite reads either listing's `--json` form or checks that `unpact` takes a directory out of both. The exit-status table is covered only at 0 and 3, both inside `scopes.sh`; status 1 for "could not answer" and status 2 for a usage error have no assertion anywhere. `.warlockignore` is not written, read or mentioned in the fixture. Neither is the `brief` → `push` → `draft` → `pull` path, which is most of what 0.1.1 does: the brief in `docs/warlock-brief-01-fix-entry-new-s-tripled-amount-add-entry-reverse-and-give.md` is the residue of a hand-run, and no script touches it. A regression in any of those surfaces is found by a person noticing, not by a FAIL line.

`check.sh` makes it worse than untested. Line 39 runs `rm -rf .warlock` before pacting, which deletes the `[[scope]]` records for `warlock-test` and `infra-ops` at lines 240-250 of the manifest and the `scope` lines they are referenced from, on `.` at line 6 and on `infra` at line 50. Running `check.sh` therefore breaks `scopes.sh`, which requires the root scope at its line 93, and breaks any later `push` or `pull` against `warlock-test`, which no longer resolves to a board. The suite that proves warlock works is the one that leaves the fixture unable to be pushed from.

The second gap is cost. `README.md` is six lines and says nothing about what running anything costs. `check.sh` pacts 16 directories and `incremental.sh` pays for a settling refresh plus one per scenario; both facts live in comment headers, so learning them means reading three scripts before running one. There is no entry point, so a release check is three commands typed from memory, in an order that matters — `scopes.sh` and `incremental.sh` both exit 2 unless `check.sh` pacted the fixture first. Someone checking a release either spends tokens they did not intend to or skips a suite and does not notice.

## Outcome

```
$ ./run.sh
warlock 0.1.1

== surfaces.sh ==
warlock 0.1.1
== the ledger lists what is pacted, and fresh is the complement of stale ==
  PASS every pacted directory is in exactly one of the two listings
  ...
== the status table a script reads ==
  PASS a listing outside a repository could not answer, and says so with 1
  PASS a malformed manifest is a 1 and not a panic
  PASS an unknown flag is a usage error, and that is a 2
== the brief front door, with nothing sent ==
  PASS a dry-run push prints the brief's title
  PASS a push at a scope with no board behind it could not answer
  PASS a malformed briefs.toml refuses before a word is sent
checks: 24 passed, 0 failed

== scopes.sh ==
warlock 0.1.1
...
checks: 31 passed, 0 failed

skipped, because they spend:
  ./check.sh        one model pass per directory, 16 directories
  ./incremental.sh  a settling refresh, then one pass per scenario
run them with ./run.sh --paid
```

`git status --porcelain` is empty afterwards, and `.warlock/pacts.toml` is byte-identical to what it was before the run.

After `./run.sh --paid`, `warlock check . --json` still reports `"scope": "warlock-test"`.

## Success criteria

**`surfaces.sh` covers the ledger listings**

- `warlock stale` and `warlock fresh` each run, and the suite asserts that every directory in `.warlock/pacts.toml` appears in exactly one of the two.
- Both are run again with `--json`, the output parses under `jq`, and the suite asserts the same membership from the parsed fields as from the plain form.
- A directory is un-pacted, and the suite asserts it leaves both listings.
- The un-pacted directory's document and the manifest are restored, and the suite asserts the manifest is byte-identical to its pre-run copy.

**`surfaces.sh` covers the exit statuses no suite reaches**

- A listing run with its working directory outside any git repository exits 1.
- A listing run against a `.warlock/pacts.toml` that is not valid TOML exits 1.
- A command given a flag warlock does not define exits 2.
- Each of the three asserts the status only, with stdout and stderr discarded.

**`surfaces.sh` covers the brief and push front door**

- `warlock push --dry-run` on `docs/warlock-brief-01-fix-entry-new-s-tripled-amount-add-entry-reverse-and-give.md` exits 0 and its output contains the brief's title line.
- `warlock push` aimed at a path whose nearest scope has no board behind it exits 1.
- With a `.warlock/briefs.toml` that is not valid TOML in place, `warlock brief </dev/null` exits 1 and `docs/` gains no file.
- No case in the section opens a socket or runs a model pass.

**Every suite leaves the checkout as it found it**

- `surfaces.sh` restores `.warlock/pacts.toml` and `.warlock/briefs.toml` from copies taken before the run, creating or removing `briefs.toml` to match what was there.
- `surfaces.sh` redirects `HOME` at a throwaway directory for the whole run, as `scopes.sh` does at line 44.
- Every restore runs from an `EXIT` trap, so it happens on a failed assertion and on an interrupt.
- `git status --porcelain` is empty after each suite and after `./run.sh`.

**`check.sh` preserves what it does not create**

- The `[[scope]]` records for `warlock-test` and `infra-ops` are byte-identical before and after `./run.sh --paid`.
- After `./run.sh --paid`, `warlock check . --json` reports `"scope": "warlock-test"` and `warlock check infra --json` reports `"scope": "infra-ops"`.
- When `check.sh` cannot put one of them back, it names on stdout which record it dropped and exits non-zero.
- `scopes.sh` passes when run immediately after `check.sh`, with no manual repair in between.

**`incremental.sh` proves `.warlockignore` suppresses a file**

- A scenario creates a file in `engine/core` and a `.warlockignore` naming it, refreshes `engine/core`, and asserts `warlock stale engine/core` is empty.
- It then edits only that file and asserts `warlock stale engine/core` is still empty, with no refresh in between.
- The scenario asserts on `engine/core` alone and never on `.`.
- `revert` removes both the created file and `.warlockignore`, and the three existing scenarios pass unchanged after it.

**`run.sh` runs the free suites and names the cost of the rest**

- `./run.sh` runs `surfaces.sh` then `scopes.sh`, and runs neither `check.sh` nor `incremental.sh`.
- `./run.sh` checks that `.warlock/pacts.toml` exists before running anything, and exits 2 with `run ./check.sh first` when it does not.
- `./run.sh --paid` runs all four with `check.sh` first, and makes no such check, because `check.sh` is what creates the manifest.
- The free run's summary names each skipped suite, what it spends, and the command that runs it.
- `./run.sh` exits non-zero when any suite it ran exited non-zero.

**The version and the costs are stated where they are needed**

- Each of `check.sh`, `scopes.sh`, `incremental.sh` and `surfaces.sh` prints the `warlock --version` output of the binary at `$WARLOCK` before its first assertion, and nothing else on that line.
- `README.md` gains a section with a row per suite: what it asserts, what it spends, and its prerequisites.
- The section states that no suite writes to Linear, and that `brief` → `push` → `draft` → `pull` is exercised by hand against the GEN board.

## Constraints

- No `lib.sh` and no shared sourced file. Each suite stays one file that reads top to bottom; the version banner is the one line duplicated across all four.
- No dependency beyond bash, git, `jq`, `tar` and coreutils. `jq` is already required by `scopes.sh`, and a suite that needs it must check for it and exit 2, as `scopes.sh` does at line 31.
- `surfaces.sh` runs no model pass and opens no socket. Every `push` in it carries `--dry-run`.
- Nothing writes to Linear.
- `.warlockignore` is never committed to this fixture. The scenario that needs one creates it and removes it.
- Assertions in `surfaces.sh` are exit statuses and `--json` fields, never prose. Prose is free to be reworded between releases and is not the contract.
- The 18 `.warlock.md` documents are left alone. They are stale after the 0.1.1 rename, including the root one that still names `pacts.toml` as the manifest and `WARLOCK.md` as the document, and `warlock refresh .` is the operator's step.
- The existing suite names and their `WARLOCK` override, `ok`/`bad` helpers and `checks: N passed` footer stay as they are. `run.sh` reads their exit statuses and nothing else.
- `check.sh` keeps pacting from nothing. Preserving the scope records means carrying them across the pact, not teaching the pact to spare them.

## Out of scope

**Status 4.** The "partly done" status needs a run that completed some directories and failed others, which means a model pass, so it cannot sit in a free suite. Provoking it with a stub binary would test the stub.

**A suite for `draft` and `pull`.** Both read and write the GEN board, and warlock 0.1.1 keeps no local record of either: the URL a push prints and the comments a draft leaves on the project are the whole record. A suite over them either leaves the board different from how it found it or asserts against a project that can be deleted out from under it. That path stays a hand exercise, and this brief is that exercise.

**Refreshing the stale documents.** Fixing all 18 costs a model pass over the whole fixture and is the operator's call, not a test. No suite asserts on their contents.

**Golden-file comparison of document text.** Settled already, and written into `check.sh`'s header: a pact reuses nothing and words every line afresh, so there is no stable text to diff. `incremental.sh` asserts byte-identity only where a line was copied rather than generated.

**Asserting the `--json` field names in this document.** The field names for `stale` and `fresh` are to be read off `warlock stale --json` on 0.1.1 while writing the suite. Naming them here would hard-code a guess.

## Scope

### 1. surfaces.sh

depends_on: []

Nothing exists yet. The harness to copy is in `scopes.sh`: `ok`/`bad` at lines 27-28, `status` at 48-53, `field` at 65-69, `checkpoint`/`unchanged` at 56-60, and the sandbox at 38-45 — `mktemp -d`, a copy of the manifest, an `EXIT` trap that puts it back, and `HOME` pointed inside the sandbox. `incremental.sh` lines 54-59 explain why the manifest has to be carried by hand: `.warlock/` is in `.gitignore`, so `git checkout` will not restore it and `git clean -fd` will not remove it. The same reasoning covers `.warlock/briefs.toml`, which may not exist at all before the run and must not exist after one that created it.

Write `surfaces.sh` as one file carrying its own copies of those helpers, with sections headed `== ... ==` in the shape `check.sh` uses, so a FAIL line still says which surface broke. The root carries `scope = "warlock-test"` and a scope covers everything beneath it until a nearer one overrides, so the suite holds that sigil inside the sandbox — `printf 'warlock-test\n' | warlock config`, as `scopes.sh` does at line 93 — or every write it attempts comes back as a 3 and the ledger section proves nothing. The ledger section runs `stale` and `fresh` plain and with `--json`, derives the set of pacted directories from the `module = "..."` lines in the manifest, and asserts each one is in exactly one listing; it then un-pacts `data` and asserts it left both. `data` rather than `infra`, because `infra` carries `scope = "infra-ops"` of its own and `scopes.sh` lines 166-171 show an un-pact across a scope the machine does not hold refuses with a 1 — a second sigil to hold for no gain. `data/.warlock.md` is tracked, so `git checkout --` restores it; the manifest comes back from the sandbox copy.

The status section runs a listing from a `mktemp -d` outside any repository for the first 1, overwrites the sandboxed manifest with text that is not TOML for the second, and passes a flag warlock does not define for the 2. Each asserts the status alone, because the prose around a status is free to be reworded between releases and the status table is the part a script reads. The brief section dry-runs a push at the committed brief and greps its output for the title line; for the refusal it must construct the no-board case itself in the sandboxed manifest, because both committed `[[scope]]` records carry `team = "GEN"` and so both resolve to a board. Exactly how a scope with no board is written is left open — read it off 0.1.1's `scope add` rather than guessing here. Then it writes a malformed `.warlock/briefs.toml` and runs `warlock brief </dev/null`, asserting the 1 and that `docs/` gained nothing: `warlock brief` reads `.warlock/brief-template.md` and `.warlock/briefs.toml` before it sends a word, so a file that will not parse is a refusal with nothing spent.

One suite rather than two, because both halves need the same sandbox discipline — a gitignored file under `.warlock/` restored by hand, `HOME` redirected so `warlock config` cannot overwrite the sigils the operator really holds — and writing that twice means getting it subtly different twice. The alternative weighed was a `lib.sh` shared by all five suites, rejected because a suite then stops being one file somebody can read top to bottom, and because `incremental.sh`'s header comments are the best documentation in this repository and would be split from the code they explain. A later edit must not move the restore out of the `EXIT` trap, must not drop the `HOME` redirect, and must not let a `push` lose its `--dry-run`: that is the line between a suite that costs seconds and one that costs money or writes to a board.

### 2. The .warlockignore scenario in incremental.sh

depends_on: []

`incremental.sh` settles with one refresh of `engine/core` at line 91, snapshots documents and manifest together in `keep_settled`, and runs three scenarios — changed file, new file, deleted file — each followed by `revert`, which is `git checkout -- .`, `git clean -qfd` and an untar of the settled snapshot. The new-file scenario at lines 112-129 already creates `settlement.rs` and relies on `git clean -qfd` to take it away, which is the pattern an ignore scenario needs.

Add a fourth scenario after the deleted-file one. It creates a file in `engine/core` and a root `.warlockignore` naming it, refreshes `engine/core`, and asserts `warlock stale engine/core` is empty. That first assertion is the precondition and not a formality: without it the second proves nothing, because a directory that was already stale would stay stale whatever the ignore said. The scenario then appends a symbol to the ignored file and asserts `warlock stale engine/core` is still empty, with no refresh between the edit and the check — the ignore is what has to suppress the file, not a refresh that re-reads it. Assertions name `engine/core` only: the root may well go stale from the new `.warlockignore` appearing under it, and that is not what the scenario tests. `revert` removes both files, so no other suite ever sees the ignore.

Creating the ignore inside the scenario rather than committing it at the root is the decision here. A committed `.warlockignore` is in force for every run, including `check.sh`'s pact, where `engine/core` is the source of the `LEDGER_VERSION`, `Posting` and `HASH_SEED` assertions — so an ignore naming anything in there can turn a passing symbol check into a FAIL that is the fixture's doing rather than warlock's. It would also shift the settled baselines the three existing scenarios compare against. A later edit must not commit `.warlockignore`, must not reorder the create-refresh-assert-edit-assert sequence, must not add a refresh between the edit and the second assertion, and must not widen the assertions to `.`.

### 3. check.sh keeps the ledger and the scopes, and run.sh fronts the lot

depends_on: [1, 2]

`check.sh` line 39 is `rm -rf .warlock`, run before the pact so that nothing is reused, and it takes the two `[[scope]]` records with it. There is no entry point. `README.md` is six lines about the fixture's shape and says nothing about cost. All four suites resolve their binary the same way — `WARLOCK="${WARLOCK:-$(command -v warlock)}"` — and none prints which version it got, so a FAIL line does not record what it was a FAIL against. `scopes.sh` and `incremental.sh` each check for a pacted fixture and exit 2 with instructions; `check.sh` is what makes that check pass.

In `check.sh`, reapply the two `[[scope]]` records after the pact and the scope lines on `.` and `infra` — `warlock scope add` with the team, review state and label each record carries, which is the shape `scopes.sh`'s `record()` helper at lines 71-73 documents. Both records are `team = "GEN"`, `review_state = "In Review"`, with label equal to the scope name. If one cannot be put back, `check.sh` names on stdout which record it dropped and exits non-zero, because a fixture that silently lost its board is the failure this slice exists to stop. The pact itself is left alone: it reuses nothing by design, and teaching it to spare files would change what `check.sh` proves.

Add one banner line to each of the four suites, after the binary check and before the first assertion, printing `"$WARLOCK" --version` and nothing else — no path, so the line stays the same wherever the binary lives. Add `run.sh`, resolving `WARLOCK` the way the suites do. Bare, it checks once that `.warlock/pacts.toml` exists, exits 2 with `run ./check.sh first` if not, then runs `surfaces.sh` and `scopes.sh`, each under its own `== name ==` header. With `--paid` it makes no such check — `check.sh` runs first and is what creates the manifest — and runs all four, `check.sh` first so the later suites have their prerequisite. It exits non-zero if any suite it ran did. The free run ends by naming each skipped suite, what it spends, and `./run.sh --paid`. Add a README section with a row per suite — what it asserts, what it spends, its prerequisites — two rows that spend nothing, `surfaces.sh` and `scopes.sh`, and two that spend model passes, `check.sh` and `incremental.sh`. `surfaces.sh` and `scopes.sh` spend nothing to run but need a pacted fixture, which only a paid `check.sh` provides, and the rows must say both. The section also states that no suite writes to Linear and that `brief` → `push` → `draft` → `pull` is exercised by hand against the GEN board.

The flag is `--paid` rather than `--all` because it names the cost, and the cost is what somebody needs warning about; `--all` names coverage, which is not the thing that surprises anyone. The alternative weighed was running everything by default behind a confirmation prompt, rejected because a prompt in a script is a thing people work around and it makes the common free case interactive. The risk in a free default is that somebody believes they checked a release when they checked half of it, which is why the summary lists the skipped suites and the exact commands rather than ending on a pass count. A later edit must not make `check.sh` or `incremental.sh` reachable without `--paid`, must not add the prerequisite check to the `--paid` path, must not drop the skipped-suite summary, and must not turn `run.sh` into the place the helpers live — each suite stays runnable on its own.
