package leetcode

import "testing"

func TestReverseBits(t *testing.T) {
	var input uint32 = 43261596   // binary: 00000010100101000001111010011100
	var answer uint32 = 964176192 // binary: 00111001011110000010100101000000

	if reverseBits(input) != answer {
		t.Errorf("error: %d != %d", input, answer)
	}
}
