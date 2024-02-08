Environment Wrapper
===================

[![Go Report Card](https://goreportcard.com/badge/github.com/ilijamt/envwrap)](https://goreportcard.com/report/github.com/ilijamt/envwrap) [![Build Status](https://travis-ci.org/ilijamt/envwrap.svg?branch=master)](https://travis-ci.org/ilijamt/envwrap) [![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/ilijamt/envwrap/blob/master/LICENSE)

A small and useful utility for tests so you can run your tests with multiple environments from inside the test functions.

The environment is set for the whole application while you use the wrapper, so running tests in parallel may have unexpected problems.

## Requirements

* [Go](https://golang.org/doc/install)

## Usage

### NewStorage

```go
package main

import (
	"fmt"
	"os"

	"github.com/ilijamt/envwrap"
)

func main() {
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
A_VARIABLE= test
A_VARIABLE=
```

### NewCleanStorage
Removes all environment variables

```go
package main

import (
	"fmt"
	"os"

	"github.com/ilijamt/envwrap"
)

func main() {
	fmt.Printf("Total items before new clean storage in environment: %d\n", len(os.Environ()))
	env := envwrap.NewCleanStorage()
	fmt.Printf("Total items in environment: %d\n", len(os.Environ()))
	_ = env.ReleaseAll()
	fmt.Printf("Total items in environment after release: %d\n", len(os.Environ()))
}

```

Should print 
```bash
Total items before new clean storage in environment: 55
Total items in environment: 0
Total items in environment after release: 55
```