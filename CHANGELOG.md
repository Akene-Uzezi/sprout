# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.13] - 2026-10-07

### Changed
- Incremented CLI version to `1.0.13`
- Deprecated the `init` command in preparation for multi-language support
- Began work on adding support for additional programming languages

## [1.0.12] - 2026-10-05

### Fixed
- Resolved bug with the CLI version display

### Changed
- Updated CLI version to `1.0.12`

## [1.0.11] - 2026-10-05

### Added
- Added version field to the CLI, enabling `sprout --version`

## [1.0.1] - 2026-10-01

### Added
- Added support for scaffolding at a specified path (in addition to the current directory)

## [1.0.0] - 2026-10-01

### Added
- Initial release
- `sprout init` command to scaffold a new Go project
- Automatic `go mod init` with provided or derived module name
- Standard project layout generation (`cmd/` and `internal/` directories)
- Starter `main.go` in `cmd/`
