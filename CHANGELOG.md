# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Changed
- Ignore client connection close errors (specifically ignoring syscall.EPIPE)

## [1.4.0] - 2020-08-29
## Added
- Include LICENSE and README.md in build archives.

## [1.3.0] - 2020-08-27
### Changed
- Renamed project to webify.

## [1.2.0] - 2020-08-27
### Added
- Started CHANGELOG.md.
- Proper help with -h.
- Options and arguments can be passed by environment variables.
- Go module files for third party deps.

### Removed
- Docker script to pass environment variable as script, as it is now supported
  directly.