# gopt

[![Go Reference](https://pkg.go.dev/badge/github.com/piotrpersona/gopt.svg)](https://pkg.go.dev/github.com/piotrpersona/gopt)

Simple golang optional package.

## Install

```sh
go get github.com/piotrpersona/gopt
```

## Usage

```go
// Create optional variable
a := gopt.Some("value")
// create none
n := gopt.None[string]()

// check if value exists
val, err := a.Get()
if err != nil {
    // it will return NoneErr
}

// or get default
val := a.Default("default")
```

