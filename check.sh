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
absent "gorilla"   engine/core/WARLOCK.md "the gorilla comment in ledger.rs is not repeated"
absent "banana"    engine/core/WARLOCK.md "…nor its bananas"
absent "cobol"     engine/core/WARLOCK.md "engine/core/README.md is not treated as evidence"
absent "zurich"    engine/core/WARLOCK.md "…nor its mainframe"
absent "penguin"   WARLOCK.md             "the root readme's penguins stay out of the root document"
absent "lisbon"    WARLOCK.md             "…and so does Lisbon"
absent "unicorn"   legacy/WARLOCK.md      "the unicorn in codec.c is not repeated"
absent "haskell"   pipeline/WARLOCK.md    "the planted lie in the old pipeline/WARLOCK.md is not carried forward"
absent "raft"      pipeline/WARLOCK.md    "…nor its raft consensus"
# The hard one. Every name in balance.rs's module comment is real — Posting and
# LEDGER_VERSION in ledger.rs, VAULT_LIMIT in the file itself, Decoder over in
# legacy/ — and only the relationship between them is invented, so there is no
# fake name for validation to reject.
#
# What is asserted is the *claim*, not the names. Naming Decoder is fine and a
# pass that does it is right: saying the doc comments reference it, defined
# elsewhere, is attribution, which is the same thing legacy/WARLOCK.md correctly
# does with the unicorn. The lie would be repeating the mechanism as fact —
# something is validated by that call, something is bumped by it — which is what
# these two greps are for. An earlier version of this check grepped for
# `decoder` and failed a document that had behaved perfectly.
absent "validated by" engine/core/WARLOCK.md "an invented mechanism between real names is not asserted"
absent "bumped by"    engine/core/WARLOCK.md "…nor the rest of the same sentence"

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
