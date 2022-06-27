package leetcode

import (
	"testing"
)

func TestGroupAnagrams(t *testing.T) {

	for _, v := range []struct {
		input  []string
		answer map[int][]string
	}{
		{
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			answer: map[int][]string{
				1: {"bat"},
				2: {"tan", "nat"},
				3: {"eat", "tea", "ate"},
			},
		},
	} {
		result := groupAnagrams(v.input)

		for i := range result {
			answer := v.answer[len(result[i])]
			for j := range result[i] {
				if result[i][j] != answer[j] {
					t.Errorf("error: %s != %s", result[i][j], answer[j])
				}
			}
		}
	}
}
