#!/usr/bin/env bash
# Check the ledger listings warlock answers about its own pacts.
#
# `warlock stale` and `warlock fresh` are the two halves of one ledger, and
# nothing else in this fixture reads either listing's `--json` form — so a
# regression there is found by somebody noticing rather than by a FAIL line.
#
# `HOME` is redirected at a throwaway directory for the whole run. Sigils live
# at `$HOME/.warlock/<name>-<digest of the canonical repo path>/config.toml`,
# and `warlock config` replaces the whole set for a repository — so without the
# redirect every run would silently overwrite whatever the operator really
# holds here.
#
# The manifest is carried by hand rather than by git. This fixture's .gitignore
# holds only `.warlock/`, so `.warlock/pacts.toml` is untracked: `git checkout`
# will not restore it and `git clean -fd` will not remove it. A copy is taken
# before anything is touched and put back from an `EXIT` trap, so a failed
# assertion and an interrupt both leave the manifest the way they found it.
#
# Assertions are exit statuses and `--json` fields, never sentences. The prose
# warlock prints is free to be reworded and is not what any of this is testing.
#
# Nothing here runs a model pass. `pact` and `refresh` never appear in a
# direction that walks a directory, so the whole script costs seconds and no
# tokens.
#
#   ./surfaces.sh
set -uo pipefail

WARLOCK="${WARLOCK:-$(command -v warlock)}"
MANIFEST=.warlock/pacts.toml
pass=0; fail=0

ok()  { printf '  \033[32mPASS\033[0m %s\n' "$1"; pass=$((pass+1)); }
bad() { printf '  \033[31mFAIL\033[0m %s\n' "$1"; fail=$((fail+1)); }

[ -x "$WARLOCK" ] || { echo "no warlock binary at $WARLOCK (set WARLOCK=)"; exit 2; }
command -v jq >/dev/null || { echo "jq is needed to read --json"; exit 2; }
[ -f "$MANIFEST" ] || { echo "no $MANIFEST — pact the fixture first"; exit 2; }

sandbox=$(mktemp -d)
cp "$MANIFEST" "$sandbox/pacts.toml.orig"
# The manifest is restored from the copy rather than from git, so a fixture with
# uncommitted pacts comes out of this the way it went in.
cleanup() { cp "$sandbox/pacts.toml.orig" "$MANIFEST"; rm -rf "$sandbox"; }
trap cleanup EXIT
export HOME="$sandbox/home"
mkdir -p "$HOME"

# status EXPECTED MESSAGE -- COMMAND...
status() {
  local want="$1" message="$2"; shift 3
  "$@" >/dev/null 2>&1
  local got=$?
  [ "$got" -eq "$want" ] && ok "$message" || bad "$message (exit $got, wanted $want)"
}

# unchanged MESSAGE — the manifest is byte-identical to the last checkpoint
checkpoint() { cp "$MANIFEST" "$sandbox/pacts.toml.mark"; }
unchanged()  {
  cmp -s "$MANIFEST" "$sandbox/pacts.toml.mark" && ok "$1" || bad "$1"
  checkpoint
}

# field PATH KEY EXPECTED MESSAGE
field() {
  local got
  got=$("$WARLOCK" check "$1" --json 2>/dev/null | jq -c ".$2")
  [ "$got" = "$3" ] && ok "$4" || bad "$4 (.$2 was $got, wanted $3)"
}

# The fixture's root carries the `warlock-test` scope, and every write below is
# made across it. The sigil is held here, inside the sandbox, where it cannot
# touch what the operator really holds for this repository.
printf 'warlock-test\n' | "$WARLOCK" config >/dev/null 2>&1 ||
  { echo "could not hold the warlock-test sigil"; exit 2; }

"$WARLOCK" --version
checkpoint

echo
printf 'checks: \033[32m%d passed\033[0m, ' "$pass"
if [ "$fail" -gt 0 ]; then printf '\033[31m%d failed\033[0m\n' "$fail"; exit 1; fi
printf '0 failed\n'
