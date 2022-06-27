package leetcode

import (
	"testing"
)

func TestClimbStairsTopDown(t *testing.T) {
	for _, v := range []struct {
		input  int
		answer int
	}{
		{input: 3, answer: 3},
		{input: 6, answer: 13},
	} {
		answer := climbStairsTopDown(v.input)
		if answer != v.answer {
			t.Errorf("error: %d != %d", answer, v.input)
		}
	}
}

func TestClimbStairsButtomUp(t *testing.T) {
	for _, v := range []struct {
		input  int
		answer int
	}{
		{input: 3, answer: 3},
		{input: 6, answer: 13},
	} {
		answer := climbStairsButtomUp(v.input)
		if answer != v.answer {
			t.Errorf("error: %d != %d", answer, v.input)
		}
	}
}
