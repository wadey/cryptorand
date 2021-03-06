// Package cryptorand provides a math/rand.Source64 implementation of crypto/rand
package cryptorand

import (
	crand "crypto/rand"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/rand"
)

type source struct{ io.Reader }

var maxInt63 = new(big.Int).SetUint64(1 << 63)
var maxUint64 = new(big.Int).Add(new(big.Int).SetUint64(math.MaxUint64), new(big.Int).SetInt64(1))

// Source is a math/rand.Source64 backed by crypto/rand.
// Calling Seed() will result in a panic.
var Source rand.Source

// NewSource returns a new rand.Source64 backed by the given random source.
// Calling Seed() will result in a panic.
func NewSource(rand io.Reader) rand.Source {
	return source{rand}
}

func init() {
	Source = NewSource(crand.Reader)
}

func (s source) Int63() int64 {
	i, err := crand.Int(s, maxInt63)
	if err != nil {
		panic(fmt.Errorf("crypto/rand.Int returned error: %v", err))
	}
	return i.Int64()
}

func (s source) Uint64() uint64 {
	i, err := crand.Int(s, maxUint64)
	if err != nil {
		panic(fmt.Errorf("crypto/rand.Int returned error: %v", err))
	}
	return i.Uint64()
}

func (source) Seed(int64) {
	panic("Seed() is not allowed on cryptorand.Source")
}
