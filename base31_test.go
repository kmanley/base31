package base31

import (
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var RAW = []uint64{0, 50, 100, 999, 1000, 1111, 5959, 99999, 123456789,
	math.MaxInt64 / 2048, math.MaxInt64 / 512, math.MaxInt64, math.MaxUint64}

var ENCODED = []string{"", "3N", "59", "339", "33A", "36V", "889", "5D3T",
	"6BQ524", "7HCE7EMPVB5", "QZCKQNEQE7H", "DR25MB7KMC659", "SF297MD67PA8H"}

func TestEncode(t *testing.T) {
	for i, v := range RAW {
		assert.Equal(t, ENCODED[i], Encode(v))
	}
}

func TestDecode(t *testing.T) {
	for i, v := range ENCODED {
		assert.Equal(t, RAW[i], Decode(v))
		assert.Equal(t, RAW[i], Decode(strings.ToLower(v)))
	}
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encode(5481594952936519619)
	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Decode("1Y2P0IJ32E8E7")
	}
}
