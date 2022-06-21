package leetcode

import (
	"testing"
)

func TestClimbStairsTopDown(t *testing.T) {
	for _, v := range []struct {
		input          int
		expectedResult int
	}{
		{input: 3, expectedResult: 3},
		{input: 6, expectedResult: 13},
	} {
		result := climbStairsTopDown(v.input)
		if result != v.expectedResult {
			t.Errorf("error: %d != %d", result, v.input)
		}
	}
}

func TestClimbStairsButtomUp(t *testing.T) {
	for _, v := range []struct {
		input          int
		expectedResult int
	}{
		{input: 3, expectedResult: 3},
		{input: 6, expectedResult: 13},
	} {
		result := climbStairsButtomUp(v.input)
		if result != v.expectedResult {
			t.Errorf("error: %d != %d", result, v.input)
		}
	}
}
