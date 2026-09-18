#!/usr/bin/env bash
# Sanity-check warlock against this fixture.
#
# Assertions are greps, never exact text. This script always runs a full pact,
# which reuses nothing and words every line afresh, so there is no golden file
# to compare against. What it checks instead is that specific strings reached a
# document, which is what distinguishes a working pass from a broken one.
#
# Every assertion here is positive, and that is the shape rather than an
# accident. The negative ones all checked that a planted falsehood had not been
# copied into a document, back when a pass was shown whatever text warlock could
# hand it. A pass is now shown code with its comments stripped, so there is
# nothing left to plant: writing a name into a file to plant it makes the name
# real. Before adding a negative check, be sure it is not that question again.
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

# present PATTERN FILE MESSAGE — the string must appear
present() {
  if [ ! -f "$2" ]; then bad "$3 (no $2)"; return; fi
  if grep -qi -- "$1" "$2"; then ok "$3"; else bad "$3"; fi
}

if [ "${1:-}" != "--no-pact" ]; then
  [ -x "$WARLOCK" ] || { echo "no warlock binary at $WARLOCK (set WARLOCK=)"; exit 2; }
  rm -rf .warlock
  find . -name WARLOCK.md -not -path ./.git/\* -delete
  git checkout -- pipeline/WARLOCK.md 2>/dev/null || true
  echo "pacting…"
  "$WARLOCK" pact . || echo "(pact reported failures — the checks below say what survived)"
  echo
fi

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
# Nothing here asserts that inventory.json is described. Two checks did, and
# both passed while its line read "contents not loaded, structure and fields
# unknown": one grepped a phrasing the pass had reworded, and the other matched
# `sku` in inventory_sku_idx over in schema.sql, a different file entirely. The
# capability they claimed was removed with the budget ladder and is not coming
# back in that shape — a line written from a sample of a 1.7 MB file is a claim
# about the part nobody read, and warlock has no way to check it, where every
# other line rests on a name witnessed in that file's own tokens. See
# ../warlock/HANDOFF.md. Do not restore them without restoring a capability.
echo "== a file nobody can read is named, not described =="
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
