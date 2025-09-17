![CI][ci-status]
[![PkgGoDev][pkg-go-dev-badge]][pkg-go-dev]

# mt

Package mt provides the [Movable Type export format][] parser.

It originally developed to parse the exported files that comes from existing blog services,
so it will not support the features that not required for importing them.

## Synopsis

See [examples][].

## Supported format version

The package supports Movable Type 8 export format.

## Currently **not** supported features

- parsing pings
- parsing custom fields
- or anything else not appears in the code

## License

see LICENSE file.

[Movable Type export format]: https://www.movabletype.jp/documentation/mt8/appendices/export-import-format/
[examples]: https://pkg.go.dev/github.com/aereal/mt#pkg-examples
[pkg-go-dev]: https://pkg.go.dev/github.com/aereal/mt
[pkg-go-dev-badge]: https://pkg.go.dev/badge/aereal/mt
[ci-status]: https://github.com/aereal/mt/workflows/ci/badge.svg?branch=main
