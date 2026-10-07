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
# assertion and an interrupt both leave the manifest the way they found it. The
# same trap carries the brief configuration `.warlock/briefs.toml` and the brief
# template beside it, which git will not restore either: whichever of them the
# run writes is removed again, because neither is in this fixture.
#
# Assertions are exit statuses and `--json` fields, never sentences. The prose
# warlock prints is free to be reworded and is not what any of this is testing.
#
# Nothing here runs a model pass. `pact` and `refresh` never appear in a
# direction that walks a directory, so the whole script costs seconds and no
# tokens.
#
# The one section that is not about the ledger is the brief front door: every
# `push` below carries `--dry-run`, so no socket is opened and nothing reaches a
# board. `push` wants a key bound to the checkout before it will say even that
# much, so a throwaway one is stored under the redirected `HOME` and goes with
# the sandbox on exit.
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

# The two files `brief` reads before it sends anything. Neither is in the
# fixture today and both sit under the gitignored `.warlock/`, so git neither
# restores nor removes them: each one's pre-run state is taken here, ahead of the
# trap, and `restore` puts back contents where there was a file and absence where
# there was none. Taking it before `trap cleanup EXIT` is what covers an
# interrupt partway through the section that writes them.
BRIEFS=.warlock/briefs.toml
TEMPLATE=.warlock/brief-template.md
for f in "$BRIEFS" "$TEMPLATE"; do
  [ -e "$f" ] && cp "$f" "$sandbox/$(basename "$f").orig"
done
restore() {
  local f saved
  for f in "$BRIEFS" "$TEMPLATE"; do
    saved="$sandbox/$(basename "$f").orig"
    if [ -e "$saved" ]; then cp "$saved" "$f"; else rm -f "$f"; fi
  done
}

# The manifest is restored from the copy rather than from git, so a fixture with
# uncommitted pacts comes out of this the way it went in.
cleanup() { cp "$sandbox/pacts.toml.orig" "$MANIFEST"; restore; rm -rf "$sandbox"; }
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

# holds NAME — this machine holds exactly NAME for this repository, and nothing
# else. `warlock config` replaces the whole set rather than adding to it, so the
# sigil the rest of this suite writes across has to be held again afterwards.
holds() { printf '%s\n' "$1" | "$WARLOCK" config >/dev/null 2>&1; }

# field PATH KEY EXPECTED MESSAGE
field() {
  local got
  got=$("$WARLOCK" check "$1" --json 2>/dev/null | jq -c ".$2")
  [ "$got" = "$3" ] && ok "$4" || bad "$4 (.$2 was $got, wanted $3)"
}

# jkey JSON FILTER EXPECTED MESSAGE — a compact jq filter over a captured answer.
# An answer that does not parse comes back empty and fails, so this is also how
# `--json` is asserted to be JSON at all.
jkey() {
  local got
  got=$(printf '%s\n' "$1" | jq -c "$2" 2>/dev/null)
  [ "$got" = "$3" ] && ok "$4" || bad "$4 ($2 was $got, wanted $3)"
}

# modules — the pacted directories, read out of the manifest's `module` lines.
# The set is derived and never written down here, so this suite covers whatever
# happens to be pacted rather than a list that goes stale beside the fixture.
modules() { sed -n 's/^module = "\(.*\)"$/\1/p' "$MANIFEST" | LC_ALL=C sort; }

# listed LISTING DIR — DIR is one whole line of LISTING
listed() { printf '%s\n' "$1" | grep -qxF "$2"; }

# has LISTING DIR MESSAGE / lacks LISTING DIR MESSAGE
has()   { listed "$1" "$2" && ok "$3" || bad "$3 ($2 is not listed)"; }
lacks() { listed "$1" "$2" && bad "$3 ($2 is still listed)" || ok "$3"; }

# partition STALE FRESH MESSAGE — every directory in `$pacted` is named by exactly
# one of the two listings. Which one it falls in is the fixture's drift and is
# never asserted; that one is in neither, or in both, is the regression.
partition() {
  local offenders="" d hits
  while IFS= read -r d; do
    hits=0
    listed "$1" "$d" && hits=$((hits+1))
    listed "$2" "$d" && hits=$((hits+1))
    [ "$hits" -eq 1 ] || offenders="$offenders $d(in $hits)"
  done <<< "$pacted"
  [ -z "$offenders" ] && ok "$3" || bad "$3 (neither or both:$offenders)"
}

# same MESSAGE A B — two listings name the same set of directories
same() {
  local a b
  a=$(printf '%s\n' "$2" | grep -v '^$' | LC_ALL=C sort)
  b=$(printf '%s\n' "$3" | grep -v '^$' | LC_ALL=C sort)
  [ "$a" = "$b" ] && ok "$1" ||
    bad "$1 (only one of them names:$(comm -3 <(printf '%s\n' "$a") <(printf '%s\n' "$b") | tr -d '\t' | tr '\n' ' '))"
}

# The fixture's root carries the `warlock-test` scope, and every write below is
# made across it. The sigil is held here, inside the sandbox, where it cannot
# touch what the operator really holds for this repository.
printf 'warlock-test\n' | "$WARLOCK" config >/dev/null 2>&1 ||
  { echo "could not hold the warlock-test sigil"; exit 2; }

"$WARLOCK" --version
checkpoint
echo

echo "== every pacted directory is in one of the two ledger listings =="
pacted=$(modules)
has "$pacted" . "the set under test comes off the manifest, root pact spelled \`.\`"
status 0 "stale answers" -- "$WARLOCK" stale
status 0 "fresh answers" -- "$WARLOCK" fresh
stale_plain=$("$WARLOCK" stale 2>/dev/null)
fresh_plain=$("$WARLOCK" fresh 2>/dev/null)
partition "$stale_plain" "$fresh_plain" "and each one is listed by exactly one of them"
unchanged "reading the ledger wrote nothing to the manifest"

echo
echo "== the --json form answers the same membership =="
stale_json=$("$WARLOCK" stale --json 2>/dev/null)
fresh_json=$("$WARLOCK" fresh --json 2>/dev/null)
jkey "$stale_json" type '"object"' "stale --json parses under jq as one object"
jkey "$fresh_json" type '"object"' "…and so does fresh --json"
jkey "$stale_json" .command '"stale"' "each names the key it answered for"
jkey "$fresh_json" .command '"fresh"' "…the other likewise"
jkey "$stale_json" '[.directories[].state] - ["stale"] | unique' '[]' \
  "no entry of the stale listing carries another state"
jkey "$fresh_json" '[.directories[].state] - ["fresh"] | unique' '[]' \
  "nor any entry of the fresh one"
stale_paths=$(printf '%s\n' "$stale_json" | jq -r '.directories[].path' 2>/dev/null)
fresh_paths=$(printf '%s\n' "$fresh_json" | jq -r '.directories[].path' 2>/dev/null)
partition "$stale_paths" "$fresh_paths" "the parsed paths partition the pacted set too"
same "stale's parsed paths are the lines it printed plain" "$stale_plain" "$stale_paths"
same "…and fresh's are the lines it printed plain" "$fresh_plain" "$fresh_paths"

echo
echo "== the status table a script reads =="
# 0 answered, 1 could not answer, 2 usage error, 3 refused having spent nothing,
# 4 partly done. A script branches on the number, so each one is worth more than
# the sentence printed beside it: a release that turned a 1 into a panic or into
# a 0 would still print something plausible, and only the status says otherwise.
# 3 is covered in scopes.sh, where a refusal can be provoked for nothing; 4 needs
# a run that walked some directories and failed others, which costs a model pass
# and so is left out of this suite deliberately.
#
# Every assertion below is the status alone — `status` already sends stdout and
# stderr to /dev/null — because the prose around a status is free to be reworded
# between releases and is not what these three are testing.

# 1, could not answer: there is no repository to answer about. The directory is
# fresh and sits under the sandbox, so nothing above it holds a `.git` and the
# `EXIT` trap carries it away with the rest. The `cd` is inside a subshell, which
# leaves this suite's own working directory where it was.
outside=$(mktemp -d "$sandbox/outside.XXXXXX")
status 1 "a listing outside any repository could not answer" -- \
  sh -c "cd '$outside' && '$WARLOCK' stale"

# 1 again, and emphatically not 101: the manifest is where it belongs but will
# not parse. A Rust panic exits 101, so asserting the exact status is what rules
# one out here. The copy goes back on the next line, before anything else reads
# the ledger — and the `EXIT` trap would put it back regardless, so a FAIL on
# this line still leaves the fixture the way it was found.
printf 'this is not a pact manifest [[[\n' > "$MANIFEST"
status 1 "a malformed manifest could not answer, and did not panic" -- "$WARLOCK" stale
cp "$sandbox/pacts.toml.orig" "$MANIFEST"

# 2, usage error: the command was never going to run at all. `--verbose` is not
# among warlock's flags, and asking for one it does not define is a usage error
# rather than a refusal or a failure to answer.
status 2 "a flag warlock does not define is a usage error" -- "$WARLOCK" stale --verbose

echo
echo "== the brief front door, with nothing sent =="
# `brief` → `push` → `draft` → `pull` is most of what warlock does, and the
# three cases here are the ones that cost nothing: a dry-run push of the brief
# this fixture already carries, a push at a scope with no board behind it, and a
# refusal to send at all because the local brief configuration will not parse.
#
# Every `push` carries `--dry-run`, which prints what would be sent and opens no
# socket. Nothing below reaches Linear and nothing below runs a model pass.
brief=docs/warlock-brief-01-fix-entry-new-s-tripled-amount-add-entry-reverse-and-give.md
[ -f "$brief" ] || { echo "no brief at $brief"; exit 2; }

# The title is read out of the brief rather than written down here, so rewording
# the document's first line does not turn into a FAIL on the line below.
title=$(sed -n '1s/^# *//p' "$brief")
[ -n "$title" ] || { echo "$brief carries no \`# \` title line to match on"; exit 2; }

# `push` refuses before anything else when no key is bound to the checkout, so a
# throwaway is stored and bound under the redirected `HOME`. It is not a Linear
# key and is never offered to one: `--dry-run` opens no socket, and the sandbox
# carries the whole store away on exit.
printf 'not-a-linear-key\n' | "$WARLOCK" key add surfaces >/dev/null 2>&1
"$WARLOCK" key use surfaces >/dev/null 2>&1

# `status` throws stdout away, so this one is captured by hand: the suite runs
# without `-e`, which is what makes `got=$?` on the next line worth reading.
out=$("$WARLOCK" push --dry-run warlock-test "$brief" 2>&1); got=$?
[ "$got" -eq 0 ] && ok "a dry-run push of the brief answers" ||
  bad "a dry-run push of the brief answers (exit $got, wanted 0)"
printf '%s\n' "$out" | grep -qF -- "$title" &&
  ok "…and names the title it read off the brief" ||
  bad "…and names the title it read off the brief (nothing in the output held \"$title\")"

# A scope held but not recorded: warlock answers 1 for it before opening a
# socket, because a name with no `[[scope]]` record is not a board here. The
# record is absent rather than removed — nothing in this suite writes one — so
# the listing below is the guard that the case is still the case it was written
# for, and not a scope somebody has since added.
scope_names=$(sed -n 's/^name = "\(.*\)"$/\1/p' "$MANIFEST" | LC_ALL=C sort)
lacks "$scope_names" no-such-board "the manifest records no scope named \`no-such-board\`"
holds "no-such-board"
status 1 "a push at a scope with no board behind it could not answer" -- \
  "$WARLOCK" push --dry-run no-such-board "$brief"

# Held sets replace rather than accumulate, so `warlock-test` comes back before
# the un-pact below, which is written across it.
holds "warlock-test" || { echo "could not hold the warlock-test sigil again"; exit 2; }

# A local brief configuration that will not parse. `brief` reads `.warlock/`
# before it sends a word, so this refusal opens no socket and runs no model pass
# — and having refused, it must not have left a brief behind in `docs/`, which is
# what the listing either side of it says. The file is absent from this fixture
# and `.gitignore` holds only `.warlock/`, so `restore` is the one thing that
# takes it away again; it runs on the next line and from the `EXIT` trap too, so
# a FAIL here still leaves `.warlock/` carrying `pacts.toml` alone.
docs_before=$(ls -A docs | LC_ALL=C sort)
printf 'this is not a brief config [[[\n' > "$BRIEFS"
status 1 "a brief config that will not parse refuses to send" -- \
  sh -c "'$WARLOCK' brief </dev/null"
restore
docs_after=$(ls -A docs | LC_ALL=C sort)
same "and the refusal wrote no brief into docs/" "$docs_before" "$docs_after"

cp "$sandbox/pacts.toml.orig" "$MANIFEST"
unchanged "and no push, sigil or refusal above wrote a line into the manifest"

echo
echo "== un-pacting takes a directory out of both listings =="
# `data` and not `infra`: infra carries a `scope` of its own that this machine
# does not hold, and an un-pact across an unheld scope refuses with 1 having done
# nothing (scopes.sh covers that). data carries no scope, so the `warlock-test`
# sigil held for the root is what opens it.
status 0 "data un-pacts across the sigil held for the root" -- "$WARLOCK" unpact data
pacted=$(modules)
lacks "$pacted" data "the manifest has stopped naming it"
stale_plain=$("$WARLOCK" stale 2>/dev/null)
fresh_plain=$("$WARLOCK" fresh 2>/dev/null)
lacks "$stale_plain" data "stale does not list it"
lacks "$fresh_plain" data "and neither does fresh"
partition "$stale_plain" "$fresh_plain" "what is still pacted is still in exactly one"
stale_paths=$("$WARLOCK" stale --json 2>/dev/null | jq -r '.directories[].path')
fresh_paths=$("$WARLOCK" fresh --json 2>/dev/null | jq -r '.directories[].path')
lacks "$stale_paths" data "the parsed stale listing has dropped it as well"
lacks "$fresh_paths" data "and so has the parsed fresh one"
partition "$stale_paths" "$fresh_paths" "with the rest partitioned as before"

echo
echo "== and the fixture is put back the way it was found =="
# The document is tracked, so git has it; the manifest is under the gitignored
# `.warlock/` and comes back from the copy taken before anything was touched.
git checkout -- data/.warlock.md
cp "$sandbox/pacts.toml.orig" "$MANIFEST"
status 0 "data/.warlock.md is what git has again" -- git diff --quiet -- data/.warlock.md
unchanged "and the manifest is byte-identical to the copy taken before the run"

echo
printf 'checks: \033[32m%d passed\033[0m, ' "$pass"
if [ "$fail" -gt 0 ]; then printf '\033[31m%d failed\033[0m\n' "$fail"; exit 1; fi
printf '0 failed\n'
