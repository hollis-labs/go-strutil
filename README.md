# go-strutil

A small, focused collection of string manipulation helpers for Go, modeled loosely on Laravel's `Illuminate\Support\Str` facade but translated to idiomatic, rune-safe Go. All functions are top-level, pure, and handle empty input gracefully. The only non-stdlib dependency is `golang.org/x/text` for Unicode normalization inside `Slugify`.

> **Status:** early internal module. Until a tagged release, consumers should pin via a `replace` directive in their own `go.mod`.

## Installation

```bash
go get github.com/hollis-labs/go-strutil
```

For local development against an uncommitted checkout:

```go
// go.mod
require github.com/hollis-labs/go-strutil v0.0.0

replace github.com/hollis-labs/go-strutil => ../../path/to/go-strutil
```

```go
import "github.com/hollis-labs/go-strutil"

slug := strutil.Slugify("Café du Monde") // "cafe-du-monde"
```

## Function reference

### Slug and normalization

| Function | Signature | Description | Example |
|---|---|---|---|
| `Slugify` | `Slugify(s string) string` | URL-safe kebab-case slug; transliterates accents via NFKD. | `Slugify("Café du Monde")` → `"cafe-du-monde"` |
| `SlugifyN` | `SlugifyN(s string, maxRunes int) string` | Slugify capped at `maxRunes` runes; trims trailing hyphens left by the cut. | `SlugifyN("hello world foo bar", 12)` → `"hello-world"` |
| `DefaultMaxSlugLength` | `const DefaultMaxSlugLength = 80` | Sensible default cap: fits `VARCHAR(100)`, keeps URLs readable. | `SlugifyN(title, strutil.DefaultMaxSlugLength)` |

### Case conversion

| Function | Signature | Description | Example |
|---|---|---|---|
| `SnakeCase` | `SnakeCase(s string) string` | Any case style → `snake_case`. | `SnakeCase("helloWorld")` → `"hello_world"` |
| `KebabCase` | `KebabCase(s string) string` | Any case style → `kebab-case` (preserves multi-byte characters). | `KebabCase("HelloWorld")` → `"hello-world"` |
| `CamelCase` | `CamelCase(s string) string` | Any case style → `camelCase`. | `CamelCase("hello_world")` → `"helloWorld"` |
| `StudlyCase` | `StudlyCase(s string) string` | Any case style → `StudlyCase` (PascalCase). | `StudlyCase("hello-world")` → `"HelloWorld"` |
| `Title` | `Title(s string) string` | Title Case — capitalizes each whitespace-separated word; underscores are *not* separators. | `Title("hello world")` → `"Hello World"` |
| `UcFirst` | `UcFirst(s string) string` | Uppercases the first rune only. | `UcFirst("hello")` → `"Hello"` |
| `LcFirst` | `LcFirst(s string) string` | Lowercases the first rune only. | `LcFirst("Hello")` → `"hello"` |

### Truncation

| Function | Signature | Description | Example |
|---|---|---|---|
| `Truncate` | `Truncate(s string, length int, suffix string) string` | Cap to `length` runes, append `suffix` if cut. | `Truncate("hello world", 5, "...")` → `"hello..."` |
| `Words` | `Words(s string, count int, suffix string) string` | Cap to `count` whitespace-separated words. | `Words("the quick brown fox", 2, "...")` → `"the quick..."` |
| `Limit` | `Limit(s string, length int) string` | Hard rune cap with no suffix. | `Limit("hello world", 5)` → `"hello"` |

### Manipulation

| Function | Signature | Description | Example |
|---|---|---|---|
| `Squish` | `Squish(s string) string` | Collapse whitespace runs to single spaces and trim. | `Squish("  a   b  ")` → `"a b"` |
| `Finish` | `Finish(s, suffix string) string` | Ensure `s` ends with `suffix`. | `Finish("dir", "/")` → `"dir/"` |
| `Start` | `Start(s, prefix string) string` | Ensure `s` starts with `prefix`. | `Start("path", "/")` → `"/path"` |
| `After` | `After(s, search string) string` | Substring after first occurrence of `search`. | `After("one.two.three", ".")` → `"two.three"` |
| `Before` | `Before(s, search string) string` | Substring before first occurrence of `search`. | `Before("one.two.three", ".")` → `"one"` |
| `Between` | `Between(s, start, end string) string` | Substring between first `start` and first subsequent `end`. | `Between("[hello]", "[", "]")` → `"hello"` |

### Inspection

| Function | Signature | Description | Example |
|---|---|---|---|
| `ContainsAll` | `ContainsAll(s string, subs []string) bool` | True iff every substring is present. | `ContainsAll("hello world", []string{"hello","world"})` → `true` |
| `ContainsAny` | `ContainsAny(s string, subs []string) bool` | True iff at least one substring is present. | `ContainsAny("hello world", []string{"mars","world"})` → `true` |

### Random

| Function | Signature | Description | Example |
|---|---|---|---|
| `Random` | `Random(length int) string` | Cryptographically secure random alphanumeric string using `crypto/rand`. | `Random(8)` → `"xK2mP9qR"` |

## Design notes

- **Rune-safe.** `Truncate`, `Limit`, and case conversions iterate runes, not bytes, so multi-byte UTF-8 strings are handled correctly.
- **Acronym splitting.** `splitWords` (internal) breaks `XMLParser` into `["XML", "Parser"]` — a run of uppercase letters followed by a lowercase letter splits before the last upper. Trailing acronyms stay intact: `parseXML` → `["parse", "XML"]`.
- **No package state.** Every function is pure; there are no init-time side effects and no global configuration.
- **No panics on bad input** (except `Random` if the OS CSPRNG is broken, which is already unrecoverable). Empty inputs yield empty outputs.
- **`Slugify` vs `KebabCase`.** `Slugify` transliterates accents and strips non-ASCII; `KebabCase` preserves multi-byte characters and only changes casing/separators.

## Development

```bash
go test ./... -v                                       # run tests
go test ./... -cover                                   # coverage (currently >96%)
go test ./... -run '^$' -fuzz='^FuzzSlugify$' -fuzztime=10s   # fuzz Slugify invariants
go test ./... -run '^$' -fuzz='^FuzzSlugifyN$' -fuzztime=10s  # fuzz SlugifyN invariants
go vet ./...                                           # static checks
gofmt -l .                                             # formatting check (should print nothing)
```

### Fuzz invariants

`FuzzSlugify` and `FuzzSlugifyN` verify the structural contract every slug must meet for any input: ASCII-only `[a-z0-9-]`, no leading/trailing hyphen, no consecutive hyphens, `Slugify` is idempotent, and `SlugifyN` always respects its rune cap.

## Out of scope (for now)

The following Laravel `Str` helpers were intentionally deferred and may land in later releases or separate modules:

- `Plural` / `Singular` — needs an inflection dictionary.
- `Markdown` / `Html` stripping — large scope.
- `Mask` / `Censor` / `Redact`.
- `Ulid` / `Uuid` — belong in their own module.
- `Is` (glob-style match).
- Fluent `Stringable` wrapper — not idiomatic Go.
