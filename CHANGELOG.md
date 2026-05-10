# Changelog

All notable changes to `go-strutil` are documented here. The format is based on
[Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this project
adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## v0.1.0 — 2026-05-10

Initial public release. Pre-1.0: API surface may evolve in minor versions;
breaking changes will be called out loudly in this changelog.

### Added

- **Slug and normalization** — `Slugify`, `SlugifyN`, `DefaultMaxSlugLength`
  constant. NFKD-based transliteration so `Café du Monde` → `cafe-du-monde`.
- **Case conversion** — `SnakeCase`, `KebabCase`, `CamelCase`, `StudlyCase`,
  `Title`, `UcFirst`, `LcFirst`. Acronym-aware tokenizer: `XMLParser` splits
  to `["XML", "Parser"]`; trailing acronyms stay intact (`parseXML` →
  `["parse", "XML"]`).
- **Truncation** — `Truncate`, `Words`, `Limit`. Rune-safe (multi-byte UTF-8
  is counted by rune, not byte).
- **Manipulation** — `Squish`, `Finish`, `Start`, `After`, `Before`, `Between`.
- **Inspection** — `ContainsAll`, `ContainsAny`.
- **Random** — `Random` (cryptographically secure alphanumeric via
  `crypto/rand`).
- Fuzz tests for `Slugify` / `SlugifyN` invariants (ASCII-only output,
  no leading/trailing or consecutive hyphens, idempotence, rune-cap respected).

### Notes

- Single non-stdlib dependency: `golang.org/x/text` (used only for Unicode
  normalization in `Slugify`).
- All functions are pure — no package-level state, no `init()` side effects.
- Empty inputs yield empty outputs across the board (no panics outside the
  documented `Random` CSPRNG-failure case).
