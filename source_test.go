package cryptorand_test

import (
	"testing"

	"github.com/wadey/cryptorand"
)

func TestSource(t *testing.T) {
	s := cryptorand.Source
	if s.Int63() == s.Int63() {
		t.Error("Expected Int63() to be random")
	}
}

func TestSeedPanics(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Error("Expected Seed() to panic")
		}
	}()
	s := cryptorand.Source
	s.Seed(1)
}

func BenchmarkRandSource(b *testing.B) {
	s := cryptorand.Source
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		s.Int63()
	}
}
