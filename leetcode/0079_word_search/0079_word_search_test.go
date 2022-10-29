package leetcode

import (
	"testing"
)

func TestExist(t *testing.T) {
	for _, v := range []struct {
		input  [][]byte
		expect bool
	}{
		{input: [][]byte{[]byte("ABCE"), []byte("SFCS"), []byte("ADEE")}, expect: true},
	} {
		result := exist(v.input, "ABCCED")
		if result != v.expect {
			t.Errorf("result %v != %v", result, v.expect)
		}
	}
}
