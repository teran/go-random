# go-random

Helper library to generate a random sequences for various purposes

## Usage example

```go
package main

import (
    "fmt"

    "github.com/teran/go-random"
)

func main() {
    fmt.Printf("New password: %s\n", random.String(random.All, 10))
}
```

## Go version warning

The library is designed to be used with Go 1.20+, usage with prior versions
if Go will cause insecure random string generation!

More detail about [math/rand](https://pkg.go.dev/math/rand#Seed) Go package
