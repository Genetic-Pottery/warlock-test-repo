#!/usr/bin/env bash
# Check the scope/sigil boundary against this fixture.
#
# `HOME` is redirected at a throwaway directory for the whole run. Sigils live
# at `$HOME/.warlock/<name>-<digest of the canonical repo path>/config.toml`,
# and `warlock config` replaces the whole set for a repository — so without the
# redirect every run would silently overwrite whatever the operator really
# holds here. The sandbox is also what lets the holds-nothing case be asserted
# at all.
#
# Assertions are exit statuses and `--json` fields, never sentences. The status
# table is the contract a script is meant to read (0 answered, 1 could not,
# 3 refused having spent nothing, 4 partly done); the prose around it is free to
# be reworded and is not what any of this is testing.
#
# Nothing here runs a model pass. `pact` and `refresh` appear only in the
# refusal direction, where the gate is asked before a directory is walked, so
# the whole script costs seconds and no tokens.
#
#   ./scopes.sh
set -uo pipefail

WARLOCK="${WARLOCK:-../warlock/target/release/warlock}"
MANIFEST=.warlock/pacts.toml
pass=0; fail=0

ok()  { printf '  \033[32mPASS\033[0m %s\n' "$1"; pass=$((pass+1)); }
bad() { printf '  \033[31mFAIL\033[0m %s\n' "$1"; fail=$((fail+1)); }

[ -x "$WARLOCK" ] || { echo "no warlock binary at $WARLOCK (set WARLOCK=)"; exit 2; }
command -v jq >/dev/null || { echo "jq is needed to read --json"; exit 2; }
[ -f "$MANIFEST" ] || { echo "no $MANIFEST — pact the fixture first"; exit 2; }
for d in legacy services services/api services/api/handlers; do
  grep -q "^module = \"$d\"\$" "$MANIFEST" ||
    { echo "$d is not pacted in $MANIFEST — run ./check.sh first"; exit 2; }
done

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

holds() { printf '%s\n' "$1" | "$WARLOCK" config >/dev/null 2>&1; }

# field PATH KEY EXPECTED MESSAGE
field() {
  local got
  got=$("$WARLOCK" check "$1" --json 2>/dev/null | jq -c ".$2")
  [ "$got" = "$3" ] && ok "$4" || bad "$4 (.$2 was $got, wanted $3)"
}

echo "== holding nothing, an unscoped path is open to anyone =="
holds ""
field . scope null "nothing scopes the root"
field . opens true "…so this machine may work there"
status 0 "a check answers even when it says no" -- "$WARLOCK" check legacy --json
status 0 "a scope can be written where nothing scopes yet" -- \
  "$WARLOCK" scope add legacy test-scope
checkpoint

echo
echo "== the scope closes the directory it is written on =="
field legacy scope '"test-scope"' "legacy is scoped"
field legacy sigils '[]' "this machine holds nothing"
field legacy opens false "so the scope is closed to it"
status 0 "the verdict is a field and never a status" -- "$WARLOCK" check legacy

echo
echo "== every mutating key refuses across it, with nothing spent =="
status 3 "scope remove is refused"       -- "$WARLOCK" scope remove legacy
status 3 "scope add is refused"          -- "$WARLOCK" scope add legacy something-else
status 3 "unpact is refused"             -- "$WARLOCK" unpact legacy
status 3 "pact is refused before it walks a directory"    -- "$WARLOCK" pact legacy
status 3 "refresh is refused before it walks a directory" -- "$WARLOCK" refresh legacy
unchanged "and the manifest is byte-identical after all five"

echo
echo "== the sigil is what opens it =="
holds "test-scope"
field legacy sigils '["test-scope"]' "the sigil is held for this repository"
field legacy opens true "and it opens the scope"
status 0 "the write that was refused now goes through" -- "$WARLOCK" scope remove legacy
checkpoint

echo
echo "== a scope covers everything beneath it until a nearer one overrides =="
status 0 "services takes a scope" -- "$WARLOCK" scope add services warlock-team
field services/api/handlers scope '"warlock-team"' "the scope reaches two directories down"
field services/api/handlers opens false "and closes it to a machine holding the wrong sigil"
holds "warlock-team"
field services/api/handlers opens true "the right sigil opens it at depth"
status 0 "a nearer scope may be written from inside an open one" -- \
  "$WARLOCK" scope add services/api inner-scope
field services/api/handlers scope '"inner-scope"' "the nearer scope is the one that applies"
field services/api/handlers opens false "…on its own, so the held outer one does not help"
field services scope '"warlock-team"' "and the outer scope still covers what is above it"
field services opens true "…where it is still held"
checkpoint

echo
echo "== a sigil is a membership test, not an expression =="
holds "billing inner-scope"
field services/api opens true "one held sigil out of several opens a matching scope"
field services opens false "an outer scope is a default below it, not a second gate"
holds "*"
field services opens true "the wildcard opens the outer scope"
field services/api opens true "…and the inner one"

echo
echo "== un-pacting refuses on blast radius, and that is a 1 and not a 3 =="
holds "warlock-team"
field services opens true "the boundary over services itself is open"
status 1 "but a scope below it that is not held refuses the un-pact" -- \
  "$WARLOCK" unpact services
unchanged "with the manifest byte-identical"

echo
echo "== a scope is a term of the pact, so un-pacting takes it away =="
holds "test-scope"
status 0 "legacy takes a scope again" -- "$WARLOCK" scope add legacy test-scope
status 0 "and the sigil un-pacts it"  -- "$WARLOCK" unpact legacy
field legacy scope null "the scope went with the pact"
field legacy opens true "leaving the path open to anyone"

echo
echo "== clearing the sigils closes every scope again =="
holds ""
field services sigils '[]' "a blank line clears what was held"
field services opens false "and a scoped path is closed to a machine holding nothing"

echo
printf 'checks: \033[32m%d passed\033[0m, ' "$pass"
if [ "$fail" -gt 0 ]; then printf '\033[31m%d failed\033[0m\n' "$fail"; exit 1; fi
printf '0 failed\n'
