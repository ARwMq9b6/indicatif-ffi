(WIP) C, Go, Python buildings to [indicatif](https://github.com/mitsuhiko/indicatif)

# Usage

Requirement: [Rust toolchain](https://www.rust-lang.org/downloads.html)

## C

Example:

```
$ make c-example-download
```

## Go

You can choose either pure go or cgo implementation

### Pure go

Requirement: [gccgo](https://gcc.gnu.org/onlinedocs/gccgo/)

Example:

```
$ go get github.com/ARwMq9b6/indicatif-ffi/indicatif-ffi-go
$ git clone git://github.com/ARwMq9b6/indicatif-ffi
$ cd indicatif-ffi
$ make go-example-download
```

### Cgo

Example:

```
$ go get github.com/ARwMq9b6/indicatif-ffi/indicatif-ffi-go
$ git clone git://github.com/ARwMq9b6/indicatif-ffi
$ cd indicatif-ffi
$ make go-cgo-example-download
```

## Python

![py3.5+](https://img.shields.io/badge/python-3.5%2B-blue.svg)

Example:

```
$ make py-example-download
```
