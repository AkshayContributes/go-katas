# go-katas

Small Go exercises, one package per kata.

## Layout

Each kata is a directory at the repo root containing a library package plus its
tests. There is no `main` — the tests are the harness.

```
go-katas/
├── go.mod
└── reverseslice/
    ├── reverseslice.go
    └── reverseslice_test.go
```

## Running

```sh
go test ./...              # all katas
go test ./reverseslice/    # one kata
go test -run TestReverseInts ./reverseslice/
go vet ./...
```

## Katas

| Kata | Problem |
| --- | --- |
| [reverseslice](./reverseslice) | Return a new slice with the elements of the input in reverse order, without mutating the input. Nil in, nil out. |

## Adding a kata

1. `mkdir <kataname>` — lowercase, no hyphens or underscores, so the directory
   name matches the package name.
2. Add `<kataname>.go` with `package <kataname>` and the exported function.
3. Add `<kataname>_test.go` with table-driven tests covering the edge cases
   (nil, empty, single element, even/odd length).
4. Add a row to the table above.
