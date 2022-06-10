# Introduction to Software Testing in GoLang 

## Testing Philosophy

Unlike other programming languages, Go's standard library doesn't come with methods to perform assertions, it relies on regular conditional statements and comparison operators to assert what's expected of an operation. INTERESTING. 

## Example

```go
// main_test.go

package main_test

import (
	"testing"
)

func TestAddition(t *testing.T) {
  got := 2 + 2
  expected := 4

  if got != expected {
    t.Errorf("Dir not get expected result. Got: '%v', wanted: '%v'", got, expected)
  }
}


```

## Types of tests in Go

- Test
  - Unit
  - Integration
  - End to End
- Benchmark
  - Performance
- Example
- Documentation

## Testing-related Packages

- `testing`
- `testing/quick`: Useful to simplify black-box testing
- `testing/iotest`: Useful to test functions that rely on I/O operations
- `net/http/testing`: Useful to test functions that perform network operations;

## Community Projects

- Testify
- Ginkgo
- GoConvey
- httpexpect
- gomock
- go-sqlmock

## Go's naming conventions

Conventions used to help Go's testing tooling to identify and run tests effectively.

- Filename must be followed by a `_test` suffix
- To blackbox the test, the package name must be followed the `_test` suffix 
- Each test function must be prefixed by a `Test` word
- Each test function expects one parameter `t *testing.T`