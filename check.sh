#!/usr/bin/env bash
# Sanity-check warlock against this fixture.
#
# Assertions are greps, never exact text. This script always runs a full pact,
# which reuses nothing and words every line afresh, so there is no golden file
# to compare against. What it checks instead is that specific strings are
# present or absent, which is what distinguishes a working pass from a broken
# one.
#
# A refresh is not the same and nothing here exercises it: it reuses a line
# wherever the recorded digest still covers both the file and the line, so
# wording does survive between runs. What it will not carry is a line somebody
# typed by hand — that fails the digest and is described again.
#
#   ./check.sh            pact the fixture, then check
#   ./check.sh --no-pact  check the documents already on disk
set -uo pipefail

WARLOCK="${WARLOCK:-../warlock/target/release/warlock}"
pass=0; fail=0

ok()   { printf '  \033[32mPASS\033[0m %s\n' "$1"; pass=$((pass+1)); }
bad()  { printf '  \033[31mFAIL\033[0m %s\n' "$1"; fail=$((fail+1)); }

# absent PATTERN FILE MESSAGE — the string must not appear
absent() {
  if [ ! -f "$2" ]; then bad "$3 (no $2)"; return; fi
  if grep -qi -- "$1" "$2"; then bad "$3"; else ok "$3"; fi
}
# present PATTERN FILE MESSAGE — the string must appear
present() {
  if [ ! -f "$2" ]; then bad "$3 (no $2)"; return; fi
  if grep -qi -- "$1" "$2"; then ok "$3"; else bad "$3"; fi
}
# An `attributed` helper stood here, passing a line that named a planted lie so
# long as it handed the claim back to the comment making it — "comment claims a
# unicorn maintains it" counted as correct behaviour. It was the wrong rule for
# what these documents are for. A line has ENTRY_CHARS to say which file to
# open, and `walk.rs` has already lost a real fact to that budget; characters
# spent attributing a joke are characters not spent routing. So the lies simply
# must not appear, whoever is credited with them.

if [ "${1:-}" != "--no-pact" ]; then
  [ -x "$WARLOCK" ] || { echo "no warlock binary at $WARLOCK (set WARLOCK=)"; exit 2; }
  rm -rf .warlock
  find . -name WARLOCK.md -not -path ./.git/\* -delete
  git checkout -- pipeline/WARLOCK.md 2>/dev/null || true
  echo "pacting…"
  "$WARLOCK" pact . || echo "(pact reported failures — the checks below say what survived)"
  echo
fi

echo "== lies that must not reach a document =="
absent "gorilla"   engine/core/WARLOCK.md "the gorilla in ledger.rs does not reach a document"
absent "banana"    engine/core/WARLOCK.md "…nor its bananas"
absent "unicorn"   legacy/WARLOCK.md      "the unicorn in codec.c does not reach a document"
absent "own taxes" legacy/WARLOCK.md      "…and the rest of the joke is not carried over"
# Six greps stood here and none of them could fail. Two were the cobol/zurich
# mainframe in engine/core/README.md, two the penguins and Lisbon in the root
# readme, two the raft-and-Haskell document planted in pipeline/WARLOCK.md.
#
# They were written when a document was summarised from whatever a directory
# held, prose included, and a lie in a README really could reach one. `own()`
# keeps only non-prose files now, which takes every `.md` out of the pass and
# the directory's own previous document with it. A README is not read, not
# listed, and not evidence; there is no path by which any of those three lies
# reaches a document, so the greps were green because the mechanism was gone.
#
# Do not add them back. A check for a failure the design excludes reads exactly
# like a check for one the model avoids, and four of this file's passes meant
# nothing for eleven days on the strength of it. The planted lies stay in the
# fixture — they cost nothing and they document what the readers ignore.

# The hard one, and the only planted lie that has ever reached a document. Every
# name in balance.rs's module comment is real — Posting and LEDGER_VERSION in
# ledger.rs, VAULT_LIMIT in the file itself, Decoder over in legacy/decoder.cpp —
# and only the relationship between them is invented, so there is no fake name
# for validation to reject.
#
# The grep is on engine/core and nowhere else, because Decoder is a real class
# and legacy/WARLOCK.md names it correctly: its own document should route a
# reader to decoder.cpp. In engine/core it routes nowhere. That is the whole
# test — a name belongs in the document of the directory that declares it.
#
# Two greps for `validated by` and `bumped by` used to stand here and caught
# only the phrasing already seen. A later `attributed` let the name through if
# the line credited the comment, which is worse: the characters are spent either
# way and the file they point at is not in this directory.
absent "decoder"   engine/core/WARLOCK.md "an invented mechanism between real names is not asserted"

echo
echo "== no document leans on a name its own directory does not declare =="
# The planted lies above can only catch what was planted, and invention is the
# one failure nobody can plant: seeding a name that is absent from the code
# would put it in the code. So this runs the other way — it derives the
# assertion from the directory instead of from a fixture, and asks of every
# name a document uses whether that directory declares it.
#
# It checks against the directory's *code*, not its bytes. Decoder appears in
# engine/core — inside the balance.rs comment that lies about it — so a check
# for the name anywhere in the directory passes the one document known to be
# wrong. A comment is not a declaration, which is the whole distinction being
# drawn, and stripping comments before the search is what draws it.
#
# The stripping is crude and stays crude. Over-stripping loses evidence and
# costs a false alarm that a reader resolves in seconds; under-stripping lets a
# comment count as a declaration and hides the defect this section exists for.
#
# There is no escape for a line that credits its source. One stood here, letting
# a name through if the line said "comment" or "elsewhere", and a document is
# not improved by spending its characters on a file this directory does not
# hold: the name still routes a reader out of the directory they are reading
# about. `## Directories` lines are the exception and are skipped below, because
# pointing at a child is what they are for.
decomment() {
  sed -e 's;//.*;;' -e 's;#.*;;' -e 's;--.*;;' \
      -e '/^[[:space:]]*\*/d' -e 's;/\*.*\*/;;' -e '/\/\*/,/\*\//d' "$@" | tr -d '\0'
}

# Identifier-shaped tokens only. Prose words are not names and must not be
# treated as claims, so a token has to carry a call, a path separator, screaming
# case or backticks to count.
names_used() {
  {
    grep -oE '[A-Za-z_][A-Za-z0-9_]*::[A-Za-z_][A-Za-z0-9_]*' <<<"$1" | tr ':' '\n'
    grep -oE '[A-Za-z_][A-Za-z0-9_]*\(' <<<"$1" | tr -d '('
    grep -oE '[A-Z][A-Z0-9]*_[A-Z0-9_]+' <<<"$1"
    grep -oE '`[A-Za-z_][A-Za-z0-9_]*`' <<<"$1" | tr -d '`'
  } 2>/dev/null | sort -u | grep -v '^$'
}

undeclared() {
  local doc="$1" dir code filenames section line stripped token found=""
  dir=$(dirname "$doc")
  local srcs=()
  while IFS= read -r f; do srcs+=("$f"); done < <(
    find "$dir" -maxdepth 1 -type f ! -name '*.md' ! -name '*.markdown' ! -name '*.mdx')
  [ "${#srcs[@]}" -eq 0 ] && { ok "$dir leans on nothing (no source files)"; return; }
  code=$(decomment "${srcs[@]}" 2>/dev/null)
  filenames=$(basename -a "${srcs[@]}" | sed 's;\..*;;')

  section=""
  while IFS= read -r line; do
    case "$line" in
      '## '*) section="$line"; continue;;
      ''|'#'*|'>'*|'<!--'*) continue;;
    esac
    # A ## Directories line describes what is below, so the names in it belong
    # to a child's code and are that child's document's problem, not this one's.
    [ "$section" = "## Directories" ] && continue
    stripped=${line#*— }
    for token in $(names_used "$stripped"); do
      [ "${#token}" -lt 3 ] && continue
      grep -qxF -- "$token" <<<"$filenames" && continue
      grep -qwF -- "$token" <<<"$code" && continue
      found="$found $token"
    done
  done < "$doc"

  found=$(tr ' ' '\n' <<<"$found" | grep -v '^$' | sort -u | tr '\n' ' ')
  if [ -n "$found" ]; then bad "$dir declares every name it uses (undeclared: $found)"
  else ok "$dir declares every name it uses"; fi
}

while IFS= read -r doc; do undeclared "$doc"; done < <(
  find . -name WARLOCK.md -not -path ./.git/\* | sort)

echo
echo "== symbols the language table must surface =="
present "LEDGER_VERSION"    engine/core/WARLOCK.md    "rust: a pub const is declared"
present "Posting"           engine/core/WARLOCK.md    "rust: a trait is declared"
present "HASH_SEED"         engine/core/WARLOCK.md    "zig: a pub const is declared"
present "NewRouter"         services/api/handlers/WARLOCK.md "go: a func is declared"
present "Boot"              services/api/WARLOCK.md   "go: a func in the parent directory is declared"
present "CART_LIMIT"        web/src/components/WARLOCK.md "typescript: an exported const is declared"
present "build_report"      tools/scripts/WARLOCK.md  "python: a def is declared"
present "REPORT_VERSION"    tools/scripts/WARLOCK.md  "python: a module const is declared"
present "Shelf"             tools/scripts/WARLOCK.md  "ruby: a class is declared"
present "Invoice"           services/billing/WARLOCK.md "java: a class is declared"
present "Stage"             pipeline/WARLOCK.md       "elixir: a defmodule is declared"

echo
echo "== no file is skipped, however big or binary =="
absent  "not read by the pass" data/WARLOCK.md "the 1.7 MB inventory.json is sampled, not skipped"
present "sku"                  data/WARLOCK.md "…and its shape reached the document"
# The property, not the phrasing. This grepped for the literal `name and size
# only`, which is one way of saying it and not the only one — a pass that wrote
# "not text" instead failed a check it had satisfied. What has to hold is that
# the line says the file is not text, rather than describing contents nobody
# read.
present "not text\|binary"     data/WARLOCK.md "logo.bin is named as not text rather than described"

echo
echo "== every directory is documented =="
for d in . engine engine/core services services/api services/api/handlers \
         services/billing web web/src web/src/components tools tools/scripts \
         pipeline infra legacy data; do
  [ -f "$d/WARLOCK.md" ] && ok "$d has a document" || bad "$d has no document"
done

echo
printf 'checks: \033[32m%d passed\033[0m, ' "$pass"
if [ "$fail" -gt 0 ]; then printf '\033[31m%d failed\033[0m\n' "$fail"; exit 1; fi
printf '0 failed\n'
