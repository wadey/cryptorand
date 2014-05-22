# cryptorand
--
    import "github.com/wadey/cryptorand"

Package cryptorand provides a math/rand.Source implementation of crypto/rand

## Usage

#### func  Source

```go
func Source() mathrand.Source
```
Source returns a math/rand.Source backed by crypto/rand. Calling Seed() will
result in a panic.
