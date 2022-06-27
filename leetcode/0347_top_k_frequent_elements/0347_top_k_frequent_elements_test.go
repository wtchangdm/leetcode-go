package leetcode

import (
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	for _, v := range []struct {
		input  []int
		topK   int
		answer []int
	}{
		{input: []int{1, 1, 2, 2, 3}, topK: 2, answer: []int{1, 2}},
		{input: []int{1}, topK: 1, answer: []int{1}},
	} {
		result := topKFrequent(v.input, v.topK)
		sort.Ints(result)
		sort.Ints(v.answer)

		for i := range result {
			if result[i] != v.answer[i] {
				t.Errorf("error: %d != %d", result[i], v.answer[i])
			}
		}
	}
}
