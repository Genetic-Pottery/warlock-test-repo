# check.sh states what warlock should do. Mostly the engine is what changes.

Work plan. Starting point is this file's commit on `main`.

**The rule this file exists to establish:** a check here states what warlock
*should* do. When a check and the engine disagree, the default assumption is
that the engine is wrong and needs changing. Weakening a check so it passes is
allowed only when the desire it encodes is itself wrong, and that has to be
argued in the commit, not assumed because fixing the engine is harder.

That rule is right and it was not being followed. But it was then over-applied:
a previous pass audited the script, found six checks that cannot fail, and wrote
all of it up as work. Most of that is a test suite testing itself. **The rule is
for checks that hide a missing capability, not a licence to spend a session on
the harness.** Two of the four sections here have been cut on that basis; what
they found is recorded below so nobody re-finds it.

Everything below was found on 2026-09-18 by running the checks against real
passes and then trying to break each one. The script reads as though it works.

## What is being done here

1. **Delete section 4's lie checks, and the planted lies themselves.** They
   tested a threat model warlock no longer has.
2. **Strengthen section 2's pair**, once the engine can satisfy it.

That is the whole of the fixture work. It is about an hour. The engine work it
depends on is in `../warlock/HANDOFF.md`, and that is where the session should
go.

## Decisions already made — do not reopen these

- **Comments do not reach a pass.** The engine strips them from the text it
  sends. A working implementation is on warlock's
  `a-comment-is-not-a-declaration` branch, unmerged; `warlock/HANDOFF.md` says
  what is in it and what has been checked.
- **A file too big to send whole is still described**, not just named. Section 2
  is engine work, not a check to weaken.

---

## 1. Delete the lie checks, and the lies

The eleven `absent` greps under `== lies that must not reach a document ==`
encode a desire: a falsehood planted in a comment, a README or a previous
document must not be copied into a `WARLOCK.md`.

**They are testing a machine that no longer exists.** They were planted back
when a document was written from whatever text warlock could hand a model, so
prose in the directory really did get picked up and repeated. That is no longer
how a document is made: `walk::own` keeps `.md` out of the file list, and the
engine strips comments from the text it sends. A pass is shown code. The
plants are not caught — they are unreachable.

So both the checks and the plants go: the unicorn in `codec.c`, the gorilla in
`ledger.rs`, the penguins in `README.md`, the mainframe in
`engine/core/README.md`. Keeping props for a play that closed is how a fixture
turns into archaeology.

**Keep the `Decoder` mechanism in `balance.rs`**, for a different reason than it
was planted for. It is the before/after case both handoffs measure against:

```
comments in : …is_settled(open) checks zero open accounts (doc claims account
              count, actually returns bool); VAULT_LIMIT = 512 applied post-decode.
comments out: Defines is_settled(open: usize) checking if open == 0, and
              VAULT_LIMIT constant set to 512.
```

It has a number attached and it is where a stripping regression would show. Do
not write a check for it here — regression cover for stripping belongs in
`cargo test`, where it runs in seconds. A 25-minute fixture run that costs model
spend is the wrong instrument for a unit-level property.

**Be careful how this is written up.** It is tempting to say no prose reaches a
pass. Not quite true: `walk::own` still collects `child_documents`, so every
child directory's `WARLOCK.md` goes to the parent's synthesis pass. Low risk —
a child's lines are checked against that child's code — but it is second-order
drift, it is not being fixed, and a commit message claiming the channel is shut
would be wrong.

Why this rather than checking the output harder, since that road was the obvious
one and was tried: it has no floor. Refused for naming a call, the pass rewrote
the same invention as prose with no identifier in it at all — `VAULT_LIMIT = 512
applied post-decode` — which no name-based check can catch. A second lie escaped
as bare ALL-CAPS nouns, which the engine's `referenced` skips on purpose because
widening it refuses most true lines. Every narrowing catches one shape and the
next answer arrives in another.

Measured cost of the input fix instead: 110 files, 2.88 MB of warlock's own
crates, 30.5% of it comments — cheaper as well as truer.

Two of these greps could never have failed anyway: `unicorn` and the rest of
that joke. A pass that would still copy a mechanism out of a comment refuses a
joke regardless. Six others were deleted for the same reason previously — see
`47bfb6d`.

**Do not add a name-checking pass to this script.** One was written and deleted.
It re-derived every name in every document and checked it against a crudely
de-commented copy of the directory — a second implementation of a rule
`document::accept_file` enforces in the engine, and the cost of two
implementations of one rule is a day spent working out which is right. It also
had a bug that silently emptied whole files: the `sed` deleted ` *` continuation
lines before running the block-comment range, so `*/` on its own line went
first, the range never closed, and every later line and file vanished. It
reported `decode` as undeclared while `decoder.cpp` declares it on line nine.

---

## 2. The over-cap file: the engine dropped a capability and the check hid it

This is the one that matters, and the clearest case of the rule at the top.

```
absent  "not read by the pass" data/WARLOCK.md "the 1.7 MB inventory.json is sampled, not skipped"
present "sku"                  data/WARLOCK.md "…and its shape reached the document"
```

Both pass. The desire they encode — stated in the section heading, *no file is
skipped, however big or binary* — is **not being met**. The document says:

```
- `inventory.json` (1.6 MB) — Large JSON inventory dataset (1.7MB); contents
  not loaded, structure and fields unknown.
```

Why each check passed anyway:

- `absent "not read by the pass"` — the pass wrote "contents not loaded,
  structure and fields unknown". Different words, same fact. The check greps
  phrasing, so any rewording escapes it.
- `present "sku"` — matched `inventory_sku_idx` in `data/schema.sql:7`. A
  different file, sent whole, in the same directory. Nothing to do with the
  JSON.

**There is no summariser in the engine.** Confirm before trusting this:

```bash
grep -rn "summar" --include='*.rs' ../warlock/crates/*/src/*.rs | grep -v tests
```

The only hits are prompt strings inside doctests. The chunked map, the demotion
ladder and the `.warlock/summaries/` cache went when per-file granularity
landed. `fitting.rs` argues for it under "No budget ladder": a directory is a
sum that need not fit, but one file either fits under `PER_FILE_BYTE_CAP` (1 MB)
or is sent as a name and a size, with nothing to demote it in favour of.
`inventory.json` is 1,755,356 bytes, so it is omitted.

That argument is about the *ladder*, which was rightly removed. It is not an
argument that a large file should go undescribed, and the section heading here
says it should not. **The decision is that it should be described.** So this is
engine work; see `warlock/HANDOFF.md` for the shape it takes, which is a head
sample rather than a restored summariser.

Expect these checks to fail until the engine is fixed. A failing check whose
desire is right is the correct intermediate state — do not quiet it.

Strengthen them so they cannot pass the way the old pair did. What the line
should carry is the file's *shape*: the fields of the records inside it. The
first 120 bytes of `inventory.json` are enough to know it is an array of
`{id, sku, qty}`, so a correct line names those.

Two traps in writing the replacement:

- **Do not grep the rendered size string** (`1.6 MB`). That couples the check to
  a formatter and breaks for a reason unrelated to the property. This was tried.
- **Check any symbol you pick against the whole directory first.** `sku` was
  already matched by `inventory_sku_idx` in `data/schema.sql`, and `qty` is
  declared there too — so neither alone proves the JSON was read. Assert on
  something only the JSON can supply, or assert on more than one field at once.

---

## The method, for when a check is written

**Break every check on purpose and confirm it goes red.** Not once at the end —
one at a time, as you write it.

```bash
cp data/WARLOCK.md /tmp/d.bak
perl -pi -e 's/\bsku\b//g' data/WARLOCK.md
bash --noprofile --norc ./check.sh --no-pact | grep -c FAIL   # expect >= 1
cp /tmp/d.bak data/WARLOCK.md
```

Every finding in this file came out of doing that. None came from reading the
script.

---

## Considered and dropped

Real findings, not being acted on. Recorded so they are not re-found and
mistaken for open work. Each is an afternoon if it ever earns one.

**The symbol checks pass on English words.** `present` greps with `-i`, and
every symbol it asks about is also an English word in some casing, so the check
is answered by prose whether or not the declared symbol reached the document.
Verified: blanking `Posting` out of `engine/core/WARLOCK.md` entirely leaves
"rust: a trait is declared" green, because the same line says "posting". Five of
eleven are like this — `Posting`, `Stage`, `Boot`, `Invoice` and `Shelf`,
answered by "posting", "stages", "boots", "invoices" and "shelf". The fix is a
`declares` helper using `grep -qwF`, with `-w` load-bearing so `Posting` is not
answered by `Postings`.

**The directory list misses a directory.** Sixteen are named in a `for` loop;
the fixture has seventeen. `monolith` (51 files) is pacted and documented on
every run and has never been checked, under a heading reading "every directory
is documented". The fix is to read the list from `.warlock/pacts.toml`, which is
what warlock actually granted, rather than from a list that drifts.

Both are true. Neither changes what warlock does for anyone.

---

## Pitfalls that cost real time

**A full `./check.sh` takes about 25 minutes.** `monolith` is 51 files and
dominates it. Do not assume it has hung.

**Never edit `check.sh` while a run is in flight.** Bash reads a script
incrementally by byte offset; inserting lines shifts everything after it and the
interpreter resumes mid-token. This corrupted a run. Wait, or stage the edit and
apply it after.

**`pgrep -f "warlock pact"` matches its own shell.** The pattern is in the
command line of the process doing the matching, so it reports the pact running
long after it exited. Read the log for progress instead.

**`grep` in a Claude Code shell is a ugrep shim** and disagrees with GNU grep on
some exit codes. Run the script as `bash --noprofile --norc ./check.sh`, and
assert on text rather than exit status.

**Background a long run with `setsid nohup … & disown`** — the harness kills
plain backgrounded jobs. A "low on memory" kill notice is not real; read
`available`, not `free`.

**`pipeline/WARLOCK.md` is a fixture input, not output.** It holds a planted
stale raft-and-Haskell document that `check.sh` restores before each pact.
Committing a regenerated document over it is what emptied four checks for eleven
days. Never `git add` it after a run. It is *not* covered by the plant deletion
in section 1 — a stale previous document is the input to a real capability,
unlike the lie plants, which have nothing left to reach.

**Restore every regenerated `WARLOCK.md` before committing**, or a document set
captured from a failing run gets enshrined:

```bash
git checkout -- $(git status --porcelain | awk '$2 ~ /WARLOCK\.md$/ {print $2}')
```

**Snapshot the documents before re-running** if you want to compare output
across two engine versions. A run deletes them all at the start, and a
comparison you did not take is gone.
