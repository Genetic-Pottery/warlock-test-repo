#!/usr/bin/env bash
# Check what a refresh reuses, against a real model.
#
# `check.sh` proves a pass. This proves the path that tries not to run one: a
# refresh keeps a file's line wherever the recorded digest still covers both the
# file and the line, and describes the rest. That reuse is a *copy*, which is
# what makes byte-identity a fair thing to assert about a model's output — the
# untouched lines are not regenerated and reworded, they are the same bytes.
#
# A second `pact` would prove none of it. A pact runs under
# `AboveFailure::Describe` and re-describes every file regardless, so the only
# entry point that reuses anything is `warlock refresh`.
#
# This one costs money: a settling refresh and then three more, each a file pass
# plus synthesis for `engine/core` and every directory above it. It needs a
# fixture that is already pacted and a clean tree, because it puts source files
# back with git — and the manifest back by hand, since `.warlock/` is ignored
# here and git will neither restore nor clean it.
#
#   ./incremental.sh
set -uo pipefail

WARLOCK="${WARLOCK:-../warlock/target/release/warlock}"
MANIFEST=.warlock/pacts.toml
SUBJECT=engine/core
DOC="$SUBJECT/WARLOCK.md"
pass=0; fail=0

ok()  { printf '  \033[32mPASS\033[0m %s\n' "$1"; pass=$((pass+1)); }
bad() { printf '  \033[31mFAIL\033[0m %s\n' "$1"; fail=$((fail+1)); }

[ -x "$WARLOCK" ] || { echo "no warlock binary at $WARLOCK (set WARLOCK=)"; exit 2; }
[ -f "$MANIFEST" ] || { echo "the fixture is not pacted — run ./check.sh first"; exit 2; }
[ -f "$DOC" ] || { echo "no $DOC — run ./check.sh first"; exit 2; }
[ -z "$(git status --porcelain)" ] || {
  echo "the tree is dirty and this script restores with git — commit or stash first"; exit 2; }

# Every file line on a page, as `name<TAB>text`. The size is dropped: adding a
# symbol to a file changes its byte count, and a sibling's size never moves, so
# keeping it would only make the changed line's own comparison noisier.
file_lines() { sed -n 's/^- `\([^`]*\)` ([^)]*) — \(.*\)$/\1\t\2/p' "$1"; }

# Every line except the named file's, which is the part a refresh must not have
# touched.
siblings() { file_lines "$1" | grep -v "^$2	" || true; }

# `.warlock/` is in this fixture's .gitignore, so the manifest is untracked:
# `git checkout` does not restore it and `git clean -fd` does not remove it. It
# has to be carried by hand, and a restore that forgets it leaves the manifest
# describing documents that were rolled back underneath it — every later
# scenario then re-describes files it should have reused, which looks exactly
# like the reuse being broken.
sandbox=$(mktemp -d)
cp "$MANIFEST" "$sandbox/pacts.toml.orig"

# Between scenarios, back to the *settled* tree rather than the committed one —
# documents and manifest together, because they are only meaningful as a pair.
keep_settled() {
  find . -name WARLOCK.md -not -path './.git/*' -print0 \
    | tar --null -cf "$sandbox/settled.tar" -T - "$MANIFEST"
}
revert() {
  git checkout -- . 2>/dev/null
  git clean -qfd
  tar xf "$sandbox/settled.tar"
}

# On the way out, back to what the checkout looked like before any of this.
trap 'git checkout -- . 2>/dev/null; git clean -qfd;
      cp "$sandbox/pacts.toml.orig" "$MANIFEST"; rm -rf "$sandbox"' EXIT

refresh() {
  if ! "$WARLOCK" refresh . >/dev/null 2>&1; then
    bad "$1 (the refresh itself failed)"
    return 1
  fi
}

# One refresh before anything is asserted, so that the documents on disk and the
# hashes in the manifest describe the same thing. A committed document can be
# behind a committed source file, and a manifest written by an older warlock
# records line hashes under a digest this one does not compute — either way the
# first refresh pays for every file once, and a baseline captured before that
# would be measuring the settling rather than the reuse.
echo "settling…"
"$WARLOCK" refresh . >/dev/null 2>&1 || { echo "the settling refresh failed"; exit 2; }
keep_settled
echo

echo "== a changed file costs its own line and no other =="
before=$(siblings "$DOC" balance.rs)
printf '\npub const SETTLEMENT_WINDOW_DAYS: u32 = 30;\n' >> "$SUBJECT/balance.rs"
if refresh "a changed file"; then
  grep -q "SETTLEMENT_WINDOW_DAYS" "$DOC" \
    && ok "the new symbol reached the document" \
    || bad "the new symbol reached the document"
  [ "$(siblings "$DOC" balance.rs)" = "$before" ] \
    && ok "every other file's line is byte-identical, because it was copied" \
    || bad "every other file's line is byte-identical, because it was copied"
  [ -z "$("$WARLOCK" stale "$SUBJECT")" ] \
    && ok "and the directory is fresh again" \
    || bad "and the directory is fresh again"
fi
revert

echo
echo "== a new file gains a line and disturbs none =="
before=$(file_lines "$DOC")
cat > "$SUBJECT/settlement.rs" <<'RS'
pub const SETTLEMENT_ATTEMPTS: u8 = 3;

pub fn settle_once() -> bool {
    true
}
RS
if refresh "a new file"; then
  grep -q "settlement.rs" "$DOC" \
    && ok "the new file has a line of its own" \
    || bad "the new file has a line of its own"
  [ "$(siblings "$DOC" settlement.rs)" = "$before" ] \
    && ok "and every line that was already there is byte-identical" \
    || bad "and every line that was already there is byte-identical"
fi
revert

echo
echo "== a deleted file's line goes with it =="
before=$(siblings "$DOC" hash.zig)
rm -f "$SUBJECT/hash.zig"
if refresh "a deleted file"; then
  grep -q '^- `hash.zig`' "$DOC" \
    && bad "the deleted file's line is off the page" \
    || ok "the deleted file's line is off the page"
  grep -q "HASH_SEED" "$DOC" \
    && bad "…and so is what it declared" \
    || ok "…and so is what it declared"
  [ "$(siblings "$DOC" hash.zig)" = "$before" ] \
    && ok "with every surviving line byte-identical" \
    || bad "with every surviving line byte-identical"
fi
revert

echo
printf 'checks: \033[32m%d passed\033[0m, ' "$pass"
if [ "$fail" -gt 0 ]; then printf '\033[31m%d failed\033[0m\n' "$fail"; exit 1; fi
printf '0 failed\n'
