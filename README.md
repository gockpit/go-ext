# [Go](https://github.com/golang/go) Extension Packages

This package provides extension APIs of [Go](https://github.com/golang/go).

Package structure follows [go/src/](https://github.com/golang/go/tree/master/src).

Following two packages are similar but provides different level of APIs.

- [goext](https://github.com/gockpit/goext) provides Go extension features.
  - Packages and APIs have almost the same with go standard packages.
- [gomore](https://github.com/gockpit/gomore) provides more extension features.
  - Additional package which are not exist in go can be provided.
  - APIs can be more higher than go standard packages.

## Naming Convention

When importing packages, use `e` prefix which means `extension`.

For example,

```go
// Import standard package.
import "slices"

// Import goext package.
import  eslices "github.com/gockpit/goext/slices"
```
