package cryptorand_test

import (
	"bytes"
	"math"
	"math/rand"
	"testing"

	"github.com/wadey/cryptorand"
)

// NOTE: We could use rand.Source64, but that would require Go1.8
type source64 interface {
	rand.Source
	Uint64() uint64
}

func TestSource(t *testing.T) {
	s := cryptorand.Source
	if s.Int63() == s.Int63() {
		t.Error("Expected Int63() to be random")
	}
}

func TestNewSource(t *testing.T) {
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
	s64 := cryptorand.NewSource(b).(source64)
	v := s64.Uint64()
	if v != 0 {
		t.Errorf("Expected Uint64() to be 0 with custom io.Reader: %v != 0", v)
	}

	b = bytes.NewReader([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})
	s64 = cryptorand.NewSource(b).(source64)
	v = s64.Uint64()
	if v != math.MaxUint64 {
		t.Errorf("Expected Uint64() to be max with custom io.Reader: %v != %v", v, uint64(math.MaxUint64))
	}

	b = bytes.NewReader([]byte{0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})
	s64 = cryptorand.NewSource(b).(source64)
	v = s64.Uint64()
	if v != math.MaxInt64 {
		t.Errorf("Unexpected Uint64() with custom io.Reader: %v != %v", v, uint64(math.MaxInt64))
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
