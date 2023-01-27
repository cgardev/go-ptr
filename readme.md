# PTR [![GoDoc][doc-img]][doc]

Transform go values to pointers and pointers to values

## Installation
`go get -u github.com/cgardev/go-ptr`


[doc-img]: https://pkg.go.dev/badge/github.com/cgardev/go-ptr
[doc]: https://pkg.go.dev/github.com/cgardev/go-ptr

## Quick Start
```go
package main

import "github.com/cgardev/go-ptr/pkg/ptr"

func main() {
	// Transform int32 to *int32
	int32Ptr := ptr.ToPtr[int32](10)
	// Transform *int32 to int32
	_ = ptr.ToValue[int32](value)
}

```