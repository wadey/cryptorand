// +build go1.9

package cryptorand_test

import (
	"bytes"
	"math"
	"math/rand"
	"testing"

	"github.com/wadey/cryptorand"
)

func TestSource64(t *testing.T) {
	s := cryptorand.Source.(rand.Source64)
	if s.Uint64() == s.Uint64() {
		t.Error("Expected Uint64() to be random")
	}
}

func TestNewSource64(t *testing.T) {
	b := bytes.NewReader(make([]byte, 8))
	s := cryptorand.NewSource(b)
	if s.Int63() != 0 {
		t.Error("Expected Int63() to be 0 with custom io.Reader")
	}

	b = bytes.NewReader([]byte{0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})
	s = cryptorand.NewSource(b)
	if s.Int63() != 1<<63-1 {
		t.Error("Expected Int63() to be max with custom io.Reader")
	}

	b = bytes.NewReader(make([]byte, 9))
	s64 := cryptorand.NewSource(b).(rand.Source64)
	v := s64.Uint64()
	if v != 0 {
		t.Errorf("Expected Uint64() to be 0 with custom io.Reader: %v != 0", v)
	}

	b = bytes.NewReader([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})
	s64 = cryptorand.NewSource(b).(rand.Source64)
	v = s64.Uint64()
	if v != math.MaxUint64 {
		t.Errorf("Expected Uint64() to be max with custom io.Reader: %v != %v", v, uint64(math.MaxUint64))
	}

	b = bytes.NewReader([]byte{0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})
	s64 = cryptorand.NewSource(b).(rand.Source64)
	v = s64.Uint64()
	if v != math.MaxInt64 {
		t.Errorf("Unexpected Uint64() with custom io.Reader: %v != %v", v, uint64(math.MaxInt64))
	}
}

func BenchmarkRandSource64(b *testing.B) {
	s := cryptorand.Source.(rand.Source64)
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		s.Uint64()
	}
}
