# go-katas

Small Go exercises, one package per kata.

## Layout

Each kata is a directory at the repo root containing a library package plus its
tests. There is no `main` — the tests are the harness.

```
go-katas/
├── go.mod
├── countwords/
│   ├── countwords.go
│   └── countwords_test.go
├── reverseslice/
│   ├── reverseslice.go
│   └── reverseslice_test.go
└── stacks/
    ├── stack.go
    └── stack_test.go
```

## Running

```sh
go test ./...              # all katas
go test ./reverseslice/    # one kata
go test -run TestReverseInts ./reverseslice/
go test -run 'TestStackPop/pop_from_empty_stack' ./stacks/   # one subtest
go test -cover ./...
go vet ./...
```

## Katas

| Kata | Problem |
| --- | --- |
| [reverseslice](./reverseslice) | Return a new slice with the elements of the input in reverse order, without mutating the input. Solved twice: a `make` + `append` loop, and `slices.Clone` + `slices.Reverse`. A nil input returns an empty, non-nil slice. |
| [countwords](./countwords) | Count the words in a string, and tally how often each one occurs. Any run of whitespace separates words, not just the space character. |
| [stacks](./stacks) | A generic LIFO stack over a doubly-linked list with a sentinel node, so `top` is never nil and `push` needs no end-case branch. `pop` and `peek` return `ErrEmptyStack` and the zero value of `T` when empty. |

## Adding a kata

1. `mkdir <kataname>` — lowercase, no hyphens or underscores, so the directory
   name matches the package name.
2. Add `<kataname>.go` with `package <kataname>` and the exported API.
3. Add `<kataname>_test.go` with table-driven tests. Cover the degenerate
   inputs first — nil, empty, single element, and whatever "one past the end"
   means for the problem. Give error cases a `wantErr` column and compare with
   `errors.Is`, so success and failure rows share one assertion.
4. Add a row to the table above.
