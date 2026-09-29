# Fix Entry::new's tripled amount, add Entry::reverse, give Router a route table

`Entry::new` in `engine/core/ledger.rs` does not store the amount it is given. It runs `for _ in 0..3 { total += amount }` and stores `total`, so every entry built through the constructor holds three times the amount the caller passed. Nothing in the repository compensates for it: the only caller is the test `an_entry_with_a_positive_amount_posts` at ledger.rs:48, which passes 5 and asserts `post` is true, so 15 satisfies it exactly as well as 5 would. The bug is invisible today and wrong for every future caller, and `post(&entry)` — the one piece of logic that reads `amount` — admits or rejects entries on a figure the caller never wrote. There is also no way to reverse an entry: `Entry` has `new` and nothing else.

`services/api/handlers/routes.go` has the matching gap on the service side. `Router` holds a `prefix` field set to `DefaultPrefix` (`/v1`) that `Handle` never reads, `NewRouter` runs a three-iteration no-op loop, and `Handle(path string) bool` returns `len(path) > 0`. There is no route table and no notion of an HTTP method, so the router answers yes to any non-empty string and `prefix` is dead weight. Its test, `TestRouterHandlesNonEmptyPaths`, asserts that a 0..25 sum equals 300 and that `/things` is handled — it checks arithmetic and a length test, not routing. Left alone, `Entry` keeps lying about amounts, there is no reversal primitive to build on, and the router cannot tell `POST /v1/entries/reverse` from `GET /v1/things`.

## Outcome

`cargo test` in `engine/` shows the ledger tests passing, including new ones: an entry built with `Entry::new(1, 5)` reports `amount == 5`; `Entry::new(1, 5).reverse(2)` yields an entry with `id == 2` and `amount == -5`; and `Entry::new(1, i64::MIN).reverse(2)` returns `None` rather than panicking.

`go test ./services/api/handlers/` passes with `TestRouterHandlesNonEmptyPaths` gone and four named tests in its place. `Handle("POST", "/v1/entries/reverse")` returns true. `Handle("GET", "/v1/entries/reverse")` returns false — the path is registered, the method is not. `Handle("POST", "/entries/reverse")` returns false, because the prefix is missing. `Handle("GET", "/v1/things")` returns false, because `/things` is not a route.

## Success criteria

**`Entry::new` stores the amount as passed**

- The `for _ in 0..3` loop and the `total` accumulator are gone from `Entry::new`.
- `Entry::new(id, amount)` constructs `Self { id, amount }`.
- A test in the existing `mod tests` asserts `Entry::new(1, 5).amount == 5`.
- `an_entry_with_a_positive_amount_posts` still passes unmodified.

**`Entry::reverse` returns a negated entry under a new id**

- `Entry` gains `pub fn reverse(&self, id: u64) -> Option<Entry>`.
- For any `amount` other than `i64::MIN`, `reverse` returns `Some(Entry)` whose `id` is the argument and whose `amount` is the arithmetic negation of `self.amount`.
- For `i64::MIN`, `reverse` returns `None`.
- `reverse` takes `&self` and does not modify it.
- A test asserts `Entry::new(1, 5).reverse(2)` gives `id == 2` and `amount == -5`.
- A test asserts `Entry::new(1, i64::MIN).reverse(2)` is `None`.

**`Router` matches method and path against a table**

- `Handle` has the signature `Handle(method, path string) bool`.
- `NewRouter` populates a route table containing exactly one entry: method `POST`, path `/entries/reverse`.
- The no-op `for i := 0; i < 3; i++` loop in `NewRouter` is gone.
- `Handle` strips `r.prefix` from `path` before matching, and returns false without consulting the table when the prefix is absent.
- Method comparison is exact and case-sensitive against the stored `POST`.
- `prefix` is read by `Handle`; no field on `Router` is written and never read.

**The handler tests name what they check**

- `TestRouterHandlesNonEmptyPaths` is deleted, including its 0..25 sum assertion.
- A test asserts `Handle("POST", "/v1/entries/reverse")` is true.
- A test asserts `Handle("GET", "/v1/entries/reverse")` is false.
- A test asserts `Handle("POST", "/entries/reverse")` is false.
- A test asserts `Handle("GET", "/v1/things")` is false.
- Each test's name states the case it covers, and none asserts arithmetic unrelated to routing.

## Constraints

- No new dependencies in either language. The Rust work uses `i64::checked_neg` from the standard library; the Go work uses the standard library only.
- `/things` must not be registered as a route. Keeping an existing test green is not a reason for a route to exist.
- `post` keeps its current signature and body. Whether a reversed entry posts follows from its sign; no reversal special case enters `post`.
- `Money` stays `i64`. No widening to `i128` and no decimal type to sidestep the `i64::MIN` case.
- `Entry`'s fields stay public and stay `id: u64, amount: Money`. Tests assert on fields directly, so no new derives on `Entry`.
- `LEDGER_VERSION`, `DEFAULT_CURRENCY`, `Posting`, `Side`, the `ledger_log!` macro and `engine/core/balance.rs` are untouched.
- The route table and its element type stay unexported. `Router`, `DefaultPrefix`, `NewRouter` and `Handle` remain the package's only exported routing surface.
- `Handle` keeps returning `bool`. It gains an argument, not a return type.

## Out of scope

- **A handler behind the route.** `Handle` reports whether a method and path match; it invokes nothing. There is no handler type, no dispatch and no call into the ledger, so `POST /v1/entries/reverse` will answer that it is a route and reverse nothing. Dispatch means designing a handler signature and a request lifecycle, a larger change than fixing a matcher, and doing it here would bury the arithmetic fix inside a routing design.
- **Any Go-to-Rust binding.** The reversal primitive is in Rust and the route is in Go with nothing between them. Wiring the two needs an FFI or transport decision this change has no basis to make.
- **Registering the rest of the API's routes.** One route proves the table matches on method and prefix. Inventing further routes with no callers would commit the service to paths nobody has asked for.
- **Path parameters, wildcards, per-route prefixes.** The table matches exact strings after one prefix strip. Patterns are worth having when a route needs them, and none does.
- **Saturating or wrapping negation.** `reverse` reports the `i64::MIN` case as `None` and leaves the decision to the caller. Silently returning `i64::MAX` or a wrapped value would substitute an amount nobody asked for — the same class of bug as the tripling.
- **Refreshing `engine/core/WARLOCK.md` and `services/api/handlers/WARLOCK.md`.** Both go stale here: the core document lists `Entry (id, amount, ::new)`, and the handlers document describes `Handle(path)` and `TestRouterHandlesNonEmptyPaths`. Regenerating them is warlock's pass over those directories, not this change's edit.

## Scope

### 1. Fix `Entry::new` and cover it

depends_on: []

Replace the body of `Entry::new` with `Self { id, amount }` and add the test asserting `Entry::new(1, 5).amount == 5`. This slice decides the constructor is a plain field assignment with no arithmetic in it, so there is nowhere for a multiplier to live again. It lands alone: the existing test keeps passing, because 5 is positive for the same reason 15 was.

### 2. Add `Entry::reverse` with the `i64::MIN` case

depends_on: [1]

Add `pub fn reverse(&self, id: u64) -> Option<Entry>` over `self.amount.checked_neg()`, with the `-5` test and the `None` test. This slice decides overflow is a value the caller sees rather than a panic or a substituted amount, which is why the return is `Option` and not `Entry`. It depends on slice 1: while `new` triples, the negation test would have to assert `-15` and would encode the bug.

### 3. Give `Router` a method-and-path route table

depends_on: []

Hold a table of method-and-path entries on `Router`, populate it in `NewRouter` with `POST /entries/reverse` in place of the no-op loop, and change `Handle` to `Handle(method, path string) bool`, stripping `r.prefix` first and returning false when the prefix is absent. This slice decides that paths are stored prefix-free and the prefix is applied at match time, so `DefaultPrefix` can change without editing route entries, and that an unregistered method on a registered path is a miss. It carries no dependency on the Rust slices, but it does not compile against the current test, so it lands with slice 4 rather than before it.

### 4. Replace the handler test with four named cases

depends_on: [3]

Delete `TestRouterHandlesNonEmptyPaths`, its 0..25 sum and its `/things` assertion, and add the four cases: the POST match, the GET method miss, the missing-prefix miss and the unregistered-path miss. This slice decides the handler tests assert routing and nothing else, and that the old test's name and arithmetic are not worth preserving under a table that makes its claim false.

Two details were left open and are the implementer's to settle: whether the table is a slice or a map, given it holds one entry; and what `Handle("GET", "/v1")` — a path that is only the prefix — should report, which falls out of how the strip is written.
