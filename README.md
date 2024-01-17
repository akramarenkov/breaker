# Breaker

[![Go Reference](https://pkg.go.dev/badge/github.com/akramarenkov/breaker.svg)](https://pkg.go.dev/github.com/akramarenkov/breaker)
[![Go Report Card](https://goreportcard.com/badge/github.com/akramarenkov/breaker)](https://goreportcard.com/report/github.com/akramarenkov/breaker)
[![codecov](https://codecov.io/gh/akramarenkov/breaker/branch/master/graph/badge.svg?token=Z8XW9Q6F2W)](https://codecov.io/gh/akramarenkov/breaker)

## Purpose

Library that allows you to break goroutine and wait it completion

## Usage

Example:

```go
package main

import (
    "fmt"

    "github.com/akramarenkov/breaker/breaker"
)

func main() {
    breaker := breaker.New()

    go func() {
        defer breaker.Complete()

        _, opened := <-breaker.IsBreaked()

        fmt.Println(opened)
    }()

    breaker.Break()

    // Output:
    // false
}
```
