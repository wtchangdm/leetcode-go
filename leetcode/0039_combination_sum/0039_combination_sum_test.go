package leetcode

import (
	"sort"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	for _, v := range []struct {
		input  []int
		target int
		expect [][]int
	}{
		{input: []int{2, 3, 6, 7}, target: 7, expect: [][]int{{7}, {2, 2, 3}}},
	} {
		result := combinationSum(v.input, v.target)

		for i := range result {
			sort.Ints(result[i])
		}
		sort.Slice(result, func(i, j int) bool {
			return result[i][0] < result[j][0]
		})

		sort.Slice(v.expect, func(i, j int) bool {
			return v.expect[i][0] < v.expect[j][0]
		})

		for i := range result {
			for j := range result[i] {
				if result[i][j] != v.expect[i][j] {
					t.Errorf("result %d != %d", result[i][j], v.expect[i][j])
				}
			}
		}
	}
}
