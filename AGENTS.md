# go-strutil

A small collection of rune-safe string helpers, modeled loosely on Laravel's
`Str` facade. Every function is top-level, pure, and handles empty input
gracefully. The only non-stdlib dependency is `golang.org/x/text`, used for
Unicode normalization inside `Slugify`.

## Start Here

- `README.md`'s function reference table is the complete API, with examples.
- `slug.go` owns `Slugify`, `SlugifyN` and `DefaultMaxSlugLength`.
- `case.go`, `manipulate.go`, `inspect.go` and `truncate.go` own the remaining
  function groups.
- `random.go` owns `Random`.
- `examples/` has a runnable program per group.

## Commands

```bash
gofmt -l .
go vet ./...
go test -race -count=1 ./...
go test -run='^$' -fuzz='^FuzzSlugify$' -fuzztime=30s
```

There is no CI workflow or Makefile in this repo. The fuzz targets
(`FuzzSlugify`, `FuzzSlugifyN`) run as ordinary tests on their seed corpus in a
normal `go test`; the last command above is only needed when changing slug
logic.

## Boundaries

`Random` draws from `crypto/rand`, not `math/rand`, and panics only when the
system entropy source is unavailable. Callers may be using it for tokens, so
swapping in `math/rand` to remove the error path or speed up a test would
quietly downgrade a security property.

`Slugify` output is a stable identifier. Applications persist slugs in URLs and
databases, so changing transliteration, the separator or the trailing-hyphen
trim rewrites keys that already exist in the wild — the fuzz targets exist to
keep that behavior pinned across edge-case input.

`DefaultMaxSlugLength` is 80 because it fits `VARCHAR(100)` with room to spare
and keeps URLs readable. It is a schema-compatibility number, not a style
preference.

Functions stay pure and top-level. There is no config object, no builder and no
package-level state, which is what makes them safe to call from anywhere.
