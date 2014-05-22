// Package cryptorand provides a math/rand.Source implementation of crypto/rand
package cryptorand

import (
	crand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
)

var max63 = new(big.Int).SetUint64(1 << 63)

type source struct{}

// Source is a math/rand.Source backed by crypto/rand.
// Calling Seed() will result in a panic.
var Source rand.Source

func init() {
	Source = source{}
}

func (source) Int63() int64 {
	i, err := crand.Int(crand.Reader, max63)
	if err != nil {
		panic(fmt.Errorf("crypto/rand.Int returned error: %v", err))
	}
	return i.Int64()
}

func (source) Seed(int64) {
	panic("Seed() is not allowed on cryptorand.Source")
}
