package utils

import (
	"fmt"
	"math"
	"strings"
)

// https://github.com/kare/base62/blob/master/base62.go
const (
	alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base     = int64(len(alphabet))
)

// Encode decoded integer to base62 string.
func (s *str) Base62Encode(n int64) string {
	if n == 0 {
		return "0"
	}

	b := make([]byte, 0, 512)
	for n > 0 {
		r := math.Mod(float64(n), float64(base))
		n /= base
		b = append([]byte{alphabet[int(r)]}, b...)
	}
	return string(b)
}

// Decode a base62 encoded string to int.
// Returns an error if input s is not valid base62 literal [0-9a-zA-Z].
func (s *str) BaseDecode(str string) (int64, error) {
	var r int64
	for _, c := range []byte(str) {
		i := strings.IndexByte(alphabet, c)
		if i < 0 {
			return 0, fmt.Errorf("unexpected character %c in base62 literal", c)
		}
		r = base*r + int64(i)
	}
	return r, nil
}