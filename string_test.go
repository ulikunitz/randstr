package randstr

import (
	"fmt"
	"testing"
)

func TestString2(t *testing.T) {
	const ab = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	t.Logf("ab = %q (%d)", ab, len(ab))
	for range 10 {
		s := String2(ab)
		t.Logf("s = %q (%d)", s, len(s))
	}
}

func BenchmarkString1(b *testing.B) {
	alphabets := []string{
		"0123456789",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
	}
	for _, ab := range alphabets {
		b.Run(fmt.Sprintf("base%d", len(ab)), func(b *testing.B) {
			for range b.N {
				String1(ab)
			}
		})
	}
}

func BenchmarkString2(b *testing.B) {
	alphabets := []string{
		"0123456789",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
	}
	for _, ab := range alphabets {
		b.Run(fmt.Sprintf("base%d", len(ab)), func(b *testing.B) {
			for range b.N {
				String1(ab)
			}
		})
	}
}

func TestString2WrongInput(t *testing.T) {
	tests := []string{
		"",
		"a",
		"aa",
		"aba",
	}
	for i, tc := range tests {
		t.Run(fmt.Sprintf("tc%d/%s", i, tc), func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Errorf("expected panic")
				}
			}()
			String2(tc)
		})
	}
}
