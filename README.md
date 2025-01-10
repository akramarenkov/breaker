# Breaker

[![Go Reference](https://pkg.go.dev/badge/github.com/akramarenkov/breaker.svg)](https://pkg.go.dev/github.com/akramarenkov/breaker)
[![Go Report Card](https://goreportcard.com/badge/github.com/akramarenkov/breaker)](https://goreportcard.com/report/github.com/akramarenkov/breaker)
[![Coverage Status](https://coveralls.io/repos/github/akramarenkov/breaker/badge.svg)](https://coveralls.io/github/akramarenkov/breaker)

## Purpose

Library that provides to break goroutine and wait it completion

## Usage

Example:

```go
package main

import (
    "fmt"

    "github.com/akramarenkov/breaker"
)

func main() {
    breaker := breaker.New()

    go func() {
        defer breaker.Complete()

        _, opened := <-breaker.IsBreaked()

        fmt.Println(opened)
    }()

    breaker.Break()

    fmt.Println(breaker.IsStopped())
    // Output:
    // false
    // true
}
```
