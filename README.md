# Sprout

A CLI tool for scaffolding new Go projects.

Sprout generates a minimal Go project layout (`cmd/`, `internal/`, and a `go.mod`)
so you can start writing code without the boilerplate.

## Install

```sh
go install go-scaffolder@latest
```

Or build from source:

```sh
git clone https://github.com/<your-username>/sprout.git
cd sprout
go build -o sprout ./cmd
```

## Usage

```sh
sprout init <path> [-m module]
```

- `<path>`: the target directory. Use `.` for the current directory or an absolute/relative path.
- `-m, --module`: the module name for `go mod init`. If omitted when scaffolding the
  current directory, Sprout derives the module name from the directory path.

### Examples

Scaffold in the current directory, deriving the module name automatically:

```sh
sprout init .
```

Scaffold at a path with an explicit module name:

```sh
sprout init ~/projects/myapp -m github.com/me/myapp
```

## What it does

Running `sprout init` will:

- Run `go mod init` with the provided (or derived) module name.
- Create a `cmd/` and `internal/` directory structure.
- Add a starter `cmd/main.go` that prints `hello world`.

## Development

This project uses [Cobra](https://github.com/spf13/cobra) for command parsing.

```sh
go build ./...
go run ./cmd init .
```

## License

See repository for license details.
