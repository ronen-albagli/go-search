# go-search

> **Learning project. Not for production use.**
> For real work use [`cloudflare/ahocorasick`](https://github.com/cloudflare/ahocorasick), `strings`, or `regexp`.

A small experiment in multi-pattern string search in Go.

## What's here

| Package | What it is |
|---|---|
| `pkg/ahocorasick` | Aho-Corasick: trie, BFS failure links, output merging |
| `pkg/binary` | Naive trie scan, O(n·m). Despite the name, no binary search. Kept as a baseline to compare against |
| `pkg/timer` | Throwaway timing helper |
| `cmd` | Demo that runs both on a sample text |

## Known limitations

- `Search` rebuilds the automaton on every call. There is no compile-once, scan-many API.
- Results are `[]string` with no match offsets.
- `ahocorasick` works on runes and `binary` on bytes, so they differ on non-ASCII text.
- Only `ahocorasick` has tests. No benchmarks.
- The timer prints to stdout from library code.

## What I was learning

- How Aho-Corasick failure links turn repeated scanning into a single pass.
- Why a trie alone is O(n·m) in the worst case.

## Run

```sh
go run ./cmd
```
