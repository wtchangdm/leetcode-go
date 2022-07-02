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
		result := climbStairsTopDown(v.input)
		if result != v.answer {
			t.Errorf("error: %d != %d", result, v.input)
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
		result := climbStairsButtomUp(v.input)
		if result != v.answer {
			t.Errorf("error: %d != %d", result, v.input)
		}
	}
}
