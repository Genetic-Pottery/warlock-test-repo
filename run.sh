#!/usr/bin/env bash
# Run the suites a release check is made of.
#
# Bare, this runs the two that cost nothing and then says what the other two
# would spend and how to ask for them. That is the whole point of the split: a
# release check used to be typed from memory in an order that mattered, and
# getting it wrong either spent tokens nobody intended or quietly skipped a
# suite.
#
#   ./run.sh         surfaces.sh, then scopes.sh — seconds, no tokens
#   ./run.sh --paid  all four, check.sh first
#
# `check.sh` goes first under --paid because it is what pacts the fixture, and
# the three after it all read what it wrote. Note that it re-pacts in place and
# so leaves the tracked documents modified; `incremental.sh` restores source
# with git and exits 2 against a dirty tree, which is why it is last rather
# than because of what it costs.
#
# No `set -e`: a suite that fails must not abort the ones after it. Each suite
# exits 1 on a FAIL and 2 on an unmet prerequisite, so the status is read after
# each one and carried to the end. Nothing else about a suite is inspected, and
# no helper is shared with one — every suite stays runnable on its own.
set -uo pipefail

WARLOCK="${WARLOCK:-$(command -v warlock)}"
MANIFEST=.warlock/pacts.toml

case "${1:-}" in
  "")     paid=0 ;;
  --paid) paid=1 ;;
  *)      echo "usage: ./run.sh [--paid]"; exit 2 ;;
esac

# Resolved the way every suite resolves it, so that one wrong WARLOCK is
# reported once here instead of four times below.
[ -x "$WARLOCK" ] || { echo "no warlock binary at $WARLOCK (set WARLOCK=)"; exit 2; }
"$WARLOCK" --version

# Both free suites exit 2 without the manifest, and `scopes.sh` wants particular
# modules named in it. Asking once, before the first header, says it once. The
# --paid path deliberately has no such check: `check.sh` runs first there, and
# the manifest is what it creates.
if [ "$paid" -eq 0 ] && [ ! -f "$MANIFEST" ]; then
  echo "the fixture is not pacted (no $MANIFEST)"
  echo "run ./check.sh first"
  exit 2
fi

status=0
suite() { # suite SCRIPT — run it under its own header and keep its status
  echo
  echo "== $1 =="
  ./"$1"; local s=$?
  if [ "$s" -ne 0 ]; then
    echo "($1 exited $s)"
    status=1
  fi
}

if [ "$paid" -eq 1 ]; then
  suite check.sh
  suite surfaces.sh
  suite scopes.sh
  suite incremental.sh
else
  suite surfaces.sh
  suite scopes.sh

  echo
  echo "skipped, because a model pass costs money:"
  echo "  check.sh        one pass per directory, over the 18 directories it pacts"
  echo "  incremental.sh  a settling refresh, then one pass per scenario, over 4 scenarios"
  echo
  echo "  ./run.sh --paid  runs all four, check.sh first"
fi

exit "$status"
