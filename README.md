Environment Wrapper
===================

[![Go Report Card](https://goreportcard.com/badge/github.com/ilijamt/envwrap)](https://goreportcard.com/report/github.com/ilijamt/envwrap) [![Build Status](https://travis-ci.org/ilijamt/envwrap.svg?branch=master)](https://travis-ci.org/ilijamt/envwrap) [![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/ilijamt/envwrap/blob/master/LICENSE)

A small and useful utility for tests so you can run your tests with multiple environments from inside the test functions.

The environment is set for the whole application while you use the wrapper, so running tests in parallel may have unexpected problems.

## Requirements

* [Go](https://golang.org/doc/install)
* [Go dep](https://github.com/golang/dep) (to install vendor deps)

## Usage

```go
import (
    "github.com/ilijamt/envwrap"
)

func main () {
    env := envwrap.NewStorage()
    defer env.ReleaseAll()
    oldVariable := os.Getenv("A_VARIABLE")
    env.Store("A_VARIABLE", "test")
    fmt.Println("ORIGINAL_VALUE=", oldVariable)
    fmt.Println("A_VARIABLE=", os.Getenv("A_VARIABLE"))
    env.Release("A_VARIABLE")
    fmt.Println("A_VARIABLE=", os.Getenv("A_VARIABLE"))
}
```

Should print 
```bash
ORIGINAL_VALUE=
A_VARIABLE=test
A_VARIABLE=
```