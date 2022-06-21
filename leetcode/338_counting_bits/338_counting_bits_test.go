package leetcode

import "testing"

func TestCountBits(t *testing.T) {
	n := 5
	bits := countBits(n)
	result := bits[len(bits)-1]
	ans := 2

	if result != ans {
		t.Errorf("error: %d != %d", result, ans)
	}
}
