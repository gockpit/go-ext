# [Gockpit](https://github.com/gockpit): [Go](https://github.com/golang/go) Extension Packages

This package provides extension APIs of [Go](https://github.com/golang/go).

Package structure follows [go/src/](https://github.com/golang/go/tree/master/src).

## Naming Convention

When importing packages, use `e` prefix which means `extension`.

For example,

```go
// Import standard package.
import "slices"

// Import go-ext package.
import  eslices "github.com/gockpit/go-ext/slices"
```
