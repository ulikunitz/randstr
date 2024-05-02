package randstr

import (
	"crypto/rand"
	"fmt"
	"math/bits"
	"unicode/utf8"
)

func _getLE64(p []byte) uint64 {
	_ = p[7]
	return uint64(p[0]) | uint64(p[1])<<8 | uint64(p[2])<<16 |
		uint64(p[3])<<24 | uint64(p[4])<<32 | uint64(p[5])<<40 |
		uint64(p[6])<<48 | uint64(p[7])<<56
}

type uint128 [2]uint64

func toUint128(p []byte) uint128 {
	var u uint128
	u[0] = _getLE64(p)
	u[1] = _getLE64(p[8:])
	return u
}

func (u *uint128) quoRem(b uint64) (r uint64) {
	u[1], r = bits.Div64(0, u[1], b)
	u[0], r = bits.Div64(r, u[0], b)
	return r
}

func stringRunes(runes []rune) string {
	p := make([]byte, 16)
	k, err := rand.Read(p)
	if err != nil {
		panic(fmt.Errorf("rand.Read: error %v", err))
	}
	if k != len(p) {
		panic("rand.Read(p) returned not enough bytes")
	}

	b := uint64(len(runes))
	l2 := bits.Len64(b - 1)
	n := (128 + l2 - 1) / l2

	r := make([]rune, n)
	x := toUint128(p)
	for i := range r {
		r[i] = runes[x.quoRem(b)]
	}
	return string(r)
}

func String1(ab string) string {
	const msgTwoRunes = "randstr: alphabet must have at least two different runes"
	if len(ab) < 2 {
		panic(msgTwoRunes)
	}
	if !utf8.ValidString(ab) {
		panic("randstr: alphabet must be a valid UTF-8 string")
	}
	runes := []rune(ab)
	if len(runes) < 2 {
		panic(msgTwoRunes)
	}
	r0 := runes[0]
	for _, r := range runes[1:] {
		if r != r0 {
			goto ok
		}
	}
	panic(msgTwoRunes)
ok:
	s := stringRunes(runes)
	return s
}

func String2(ab string) string {
	const msgTwoRunes = "randstr: alphabet must have at least two different runes"
	if len(ab) < 2 {
		panic(msgTwoRunes)
	}
	if !utf8.ValidString(ab) {
		panic("randstr: alphabet must be a valid UTF-8 string")
	}
	runes := []rune(ab)
	if len(runes) < 2 {
		panic(msgTwoRunes)
	}
	for i, r := range runes {
		for _, t := range runes[:i] {
			if r == t {
				panic("randstr: all runes in the alphabet must be unique")
			}
		}
	}
	s := stringRunes(runes)
	return s
}
