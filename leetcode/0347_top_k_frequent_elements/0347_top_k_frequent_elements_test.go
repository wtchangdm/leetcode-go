package leetcode

import (
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	for _, v := range []struct {
		input  []int
		topK   int
		result []int
	}{
		{input: []int{1, 1, 2, 2, 3}, topK: 2, result: []int{1, 2}},
		{input: []int{1}, topK: 1, result: []int{1}},
	} {
		answer := topKFrequent(v.input, v.topK)
		sort.Ints(answer)
		sort.Ints(v.result)

		for i := range answer {
			if answer[i] != v.result[i] {
				t.Errorf("error: %d != %d", answer[i], v.result[i])
			}
		}
	}
}
