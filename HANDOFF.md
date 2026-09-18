# check.sh is the specification. The engine is what has to fit it.

Work plan for a fresh agent. Starting point is this file's commit on `main`.

**The rule this file exists to establish:** a check here states what warlock
*should* do. When a check and the engine disagree, the default assumption is
that the engine is wrong and needs changing. Weakening a check so it passes is
allowed only when the desire it encodes is itself wrong, and that has to be
argued in the commit, not assumed because fixing the engine is harder.

That rule was not being followed. Everything below was found on 2026-09-18 by
running the checks against real passes and then trying to break each one.
Nothing here came from reading; the script reads as though it works.

There are two kinds of defect, and they need opposite fixes:

- **A check that cannot fail.** It reports coverage the suite does not have.
  The check is wrong — it fails to express its own desire. Fix the check.
- **A check that passes while the desire is unmet.** Worse, because it hides a
  capability the engine has quietly lost. Fix the engine.

Six checks cannot fail. Two more are the second kind. Do these in order and
prove each one can fail before moving on.

---

## The method, which matters more than any single change

**Break every check on purpose and confirm it goes red.** Not once at the end —
one at a time, as you write it.

```bash
cp data/WARLOCK.md /tmp/d.bak
perl -pi -e 's/\bsku\b//g' data/WARLOCK.md
bash --noprofile --norc ./check.sh --no-pact | grep -c FAIL   # expect >= 1
cp /tmp/d.bak data/WARLOCK.md
```

Every finding below came out of doing that. None came from reading the script.

---

## 1. The symbol checks pass on English words — fix the check

`present` greps with `-i`, and every symbol it asks about is also an English
word in some casing. The check is answered by prose whether or not the declared
symbol reached the document.

Verified: blanking `Posting` out of `engine/core/WARLOCK.md` entirely leaves
"rust: a trait is declared" **green**, because the same line says "posting".
Five of the eleven are like this — `Posting`, `Stage`, `Boot`, `Invoice` and
`Shelf`, answered by "posting", "stages", "boots", "invoices" and "shelf".

The desire is right; the check does not express it. Add a third helper beside
`absent` and `present`:

```bash
# declares SYMBOL FILE MESSAGE — the identifier must appear, spelt exactly and whole
declares() {
  if [ ! -f "$2" ]; then bad "$3 (no $2)"; return; fi
  if grep -qwF -- "$1" "$2"; then ok "$3"; else bad "$3"; fi
}
```

`-w` is the load-bearing flag: `Posting` must not be answered by `Postings`,
and a name in backticks is still whole because a backtick is not a word
character. `-F` stops a symbol being read as a regex.

Change all eleven `present` lines under `== symbols the language table must
surface ==` to `declares`. Leave the other `present` calls alone — those are
about what a line *says*, not which identifier it spells.

**Verify:** remove each of the eleven symbols in turn with
`perl -pi -e 's/\bSYM\b//g'` and confirm each produces a failure. All eleven
must. Before the change, five do not.

---

## 2. The over-cap file: the engine dropped a capability and the check hid it

This is the important one, and the clearest case of the rule at the top.

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

The only hits are prompt strings inside doctests. The chunked map, the
demotion ladder and the `.warlock/summaries/` cache went when per-file
granularity landed. `fitting.rs` argues for it under "No budget ladder": a
directory is a sum that need not fit, but one file either fits under
`PER_FILE_BYTE_CAP` (1 MB) or is sent as a name and a size, with nothing to
demote it in favour of. `inventory.json` is 1,755,356 bytes, so it is omitted.

That argument is about the *ladder*, which was rightly removed. It is not an
argument that a large file should go undescribed, and the section heading here
says it should not. So this is engine work, and there is a decision to make
first:

- **If a big file should still be described** — restore a summariser for the
  over-cap case. Then strengthen these checks so they cannot pass on phrasing
  or on a neighbouring file, and expect them red until the engine is fixed. A
  red check whose desire is right is the correct intermediate state; do not
  quiet it.
- **If a big file should now just be named and sized** — then the desire has
  changed, the heading is stale, and the checks should be rewritten to assert
  that instead (`declares "inventory.json" data/WARLOCK.md`). Argue it in the
  commit message. Do not do this silently because it is the easier road.

Whichever way it goes, the replacement check must not grep the rendered size
string (`1.6 MB`) — that couples it to a formatter and breaks for a reason
unrelated to the property. A stronger idea was tried and rejected: a field only
the unsent file holds would prove invention if it appeared, but the candidate
`qty` is declared in `data/schema.sql` too. Check any candidate against the
whole directory first.

---

## 3. The directory list misses a directory — fix the check

```bash
for d in . engine engine/core services … data; do
```

Sixteen directories listed; the fixture has seventeen. `monolith` (51 files) is
pacted and documented on every run and has never been checked, under a heading
reading "every directory is documented".

A list that names its own subjects drifts from what it checks, which is the
same failure as a stale comment. Read it from the manifest, which is what
warlock granted:

```bash
if [ ! -f .warlock/pacts.toml ]; then
  bad "no manifest to read the pacted directories from"
else
  pacted=$(sed -n 's/^module = "\(.*\)"$/\1/p' .warlock/pacts.toml)
  [ -n "$pacted" ] || bad "the manifest records no pacted directory at all"
  while IFS= read -r d; do
    [ -n "$d" ] || continue
    [ -f "${d%/}/WARLOCK.md" ] && ok "$d has a document" || bad "$d has no document"
  done <<<"$pacted"
fi
```

**Verify:** `monolith` appears in the output; moving any one document aside
fails; moving `.warlock/pacts.toml` aside fails rather than silently checking
nothing.

---

## 4. The lie checks — decide what the desired truth is

The eleven `absent` greps under `== lies that must not reach a document ==`
encode a desire: a falsehood planted in a comment, a README or a previous
document must not be copied into a WARLOCK.md.

Under the rule at the top, the question is not whether the engine currently
satisfies them. It is whether the desire is still the one you want. There are
two coherent answers and they lead to different engines:

- **A pass may read comments and must not repeat their claims.** Then these
  checks stay, and the engine needs a check on what a pass writes. Be warned
  that this was attempted and has no floor: refused for a call, the pass
  rewrote the same invention as prose with no identifier in it at all
  (`VAULT_LIMIT = 512 applied post-decode`), and a second lie escaped as bare
  ALL-CAPS nouns, which the engine's `referenced` skips on purpose because
  widening it refuses most true lines. Every narrowing catches one shape.
- **A pass is shown code and nothing else.** Then the falsehood never arrives,
  these checks become untestable by construction — planting a name in a file
  makes the name real — and the section goes. Measured cost of this on
  warlock's own crates: 110 files, 2.88 MB, 30.5% of it comments, so it is also
  cheaper. Measured effect on the line for `balance.rs`:

  ```
  comments in : …is_settled(open) checks zero open accounts (doc claims account
                count, actually returns bool); VAULT_LIMIT = 512 applied post-decode.
  comments out: Defines is_settled(open: usize) checking if open == 0, and
                VAULT_LIMIT constant set to 512.
  ```

Two of these greps cannot fail under either answer and should go regardless:
`unicorn` and the rest of that joke. A pass that would still copy a mechanism
out of a comment refuses a joke anyway. Six others were deleted for the same
reason previously — see `47bfb6d`.

If the section goes, leave the planted text in the fixture files — the unicorn
in `codec.c`, the gorilla in `ledger.rs`, the penguins in `README.md`, the
mainframe in `engine/core/README.md`, the `Decoder` mechanism in `balance.rs`.
It costs nothing and keeps a directory holding a plausible falsehood in the set
a run walks over. Do not write a check for one.

**Do not add a name-checking pass to this script.** One was written and
deleted. It re-derived every name in every document and checked it against a
crudely de-commented copy of the directory — a second implementation of a rule
`document::accept_file` enforces in the engine, and the cost of two
implementations of one rule is a day spent working out which is right. It also
had a bug that silently emptied whole files: the `sed` deleted ` *`
continuation lines before running the block-comment range, so `*/` on its own
line went first, the range never closed, and every later line and file
vanished. It reported `decode` as undeclared while `decoder.cpp` declares it on
line nine.

---

## Pitfalls that cost real time

**A full `./check.sh` takes about 25 minutes.** `monolith` is 51 files and
dominates it. Do not assume it has hung.

**Never edit `check.sh` while a run is in flight.** Bash reads a script
incrementally by byte offset; inserting lines shifts everything after it and
the interpreter resumes mid-token. This corrupted a run. Wait, or stage the
edit and apply it after.

**`pgrep -f "warlock pact"` matches its own shell.** The pattern is in the
command line of the process doing the matching, so it reports the pact running
long after it exited. Read the log for progress instead.

**`grep` in a Claude Code shell is a ugrep shim** and disagrees with GNU grep
on some exit codes. Run the script as `bash --noprofile --norc ./check.sh`, and
assert on text rather than exit status.

**Background a long run with `setsid nohup … & disown`** — the harness kills
plain backgrounded jobs. A "low on memory" kill notice is not real; read
`available`, not `free`.

**`pipeline/WARLOCK.md` is a fixture input, not output.** It holds a planted
stale raft-and-Haskell document that `check.sh` restores before each pact.
Committing a regenerated document over it is what emptied four checks for
eleven days. Never `git add` it after a run.

**Restore every regenerated `WARLOCK.md` before committing**, or a document set
captured from a failing run gets enshrined:

```bash
git checkout -- $(git status --porcelain | awk '$2 ~ /WARLOCK\.md$/ {print $2}')
```

**Snapshot the documents before re-running** if you want to compare output
across two engine versions. A run deletes them all at the start, and a
comparison you did not take is gone.

---

## What "done" looks like

Every check in the file has been made to fail on purpose. Any check still red
is red because the engine owes it work, and that is written down somewhere with
the engine change it is waiting for — not quietly reworded until it passes.
