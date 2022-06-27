package leetcode

import "testing"

func TestCountBits(t *testing.T) {
	n := 5
	bits := countBits(n)
	answer := bits[len(bits)-1]
	ans := 2

	if answer != ans {
		t.Errorf("error: %d != %d", answer, ans)
	}
}
