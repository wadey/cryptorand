# cryptorand
--
    import "github.com/wadey/cryptorand"

Package cryptorand provides a math/rand.Source implementation of crypto/rand

## Usage

```go
var Source rand.Source
```
Source is a math/rand.Source backed by crypto/rand. Calling Seed() will result
in a panic.
