package base31

import "math/big"

var (
	base31 = [31]byte{
		'2', '3', '4', '5', '6', '7', '8', '9',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H',
		'J', 'K', 'M', 'N', 'P', 'Q', 'R', 'S',
		'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

	index = map[byte]int{
		'2': 0, '3': 1, '4': 2, '5': 3, '6': 4,
		'7': 5, '8': 6, '9': 7, 'A': 8, 'B': 9,
		'C': 10, 'D': 11, 'E': 12, 'F': 13, 'G': 14,
		'H': 15, 'J': 16, 'K': 17, 'M': 18, 'N': 19,
		'P': 20, 'Q': 21, 'R': 22, 'S': 23, 'T': 24,
		'U': 25, 'V': 26, 'W': 27, 'X': 28, 'Y': 29,
		'Z': 30, 'a': 8, 'b': 9, 'c': 10, 'd': 11,
		'e': 12, 'f': 13, 'g': 14, 'h': 15, 'j': 16,
		'k': 17, 'm': 18, 'n': 19, 'p': 20, 'q': 21,
		'r': 22, 's': 23, 't': 24, 'u': 25, 'v': 26,
		'w': 27, 'x': 28, 'y': 29, 'z': 30}
)

// Encode encodes a value to base31
func Encode(value uint64) string {
	var res [16]byte
	var i int
	for i = len(res) - 1; value != 0; i-- {
		res[i] = base31[value%31]
		value /= 31
	}
	return string(res[i+1:])
}

// Decode decodes a base31-encoded string
/* TODO: NOTE: this worked fine for all test cases except
   math.MaxInt64, math.MaxUint64, rather than spend time
   trying to find a workaround for the floating point rounding
   I replaced with Pow with math/big Exp which is slower but precise.
func Decode(s string) uint64 {
	res := uint64(0)
	l := len(s) - 1
	for idx := range s {
		c := s[l-idx]
		byteOffset := index[c]
		res += uint64(byteOffset) * uint64(math.Pow(31, float64(idx)))
	}
	return res
}
*/

func Decode(s string) uint64 {
	res := uint64(0)
	l := len(s) - 1
	b31 := big.NewInt(31)
	bidx := big.NewInt(0)
	bpow := big.NewInt(0)
	for idx := range s {
		c := s[l-idx]
		byteOffset := index[c]
		bidx.SetUint64(uint64(idx))
		res += uint64(byteOffset) * bpow.Exp(b31, bidx, nil).Uint64()
	}
	return res
}
