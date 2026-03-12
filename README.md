# ximg

`ximg` is a Go CLI for deep container image comparison.

It is meant to compare two container images and explain:

- package differences
- size differences
- missing and extra content
- later, dependency reasons for why a package exists

## What is implemented

- a Go module
- a minimal CLI entrypoint
- a `diff` command stub
- a shared report type
- a basic test
- beginner-friendly notes in `docs/`

Current commands:

```bash
ximg version
ximg diff <image-a> <image-b>
```

Right now, `diff` only validates input and prints a placeholder result.

## How to use

```bash
go run ./cmd/ximg version
go run ./cmd/ximg diff image-a image-b
```

To run this, you need Go installed first.

Useful commands:

```bash
go test ./...
go build -o bin/ximg ./cmd/ximg
```

Learning notes for this project are in:

- [docs/00-learning-path.md](/Users/kshitijagarwal/projects/ximg/docs/00-learning-path.md)
- [docs/01-bootstrap-cli.md](/Users/kshitijagarwal/projects/ximg/docs/01-bootstrap-cli.md)

## Future plans

1. Add image loading and metadata inspection.
2. Add distro detection and installed package scanning.
3. Add package diffing.
4. Add package dependency ancestry reasoning.
5. Add filesystem size analysis.
6. Merge all of that into a useful explanation report.
