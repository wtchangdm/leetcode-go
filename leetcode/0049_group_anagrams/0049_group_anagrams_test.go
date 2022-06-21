package leetcode

import (
	"testing"
)

func TestGroupAnagrams(t *testing.T) {

	for _, v := range []struct {
		input  []string
		result map[int][]string
	}{
		{
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			result: map[int][]string{
				1: {"bat"},
				2: {"tan", "nat"},
				3: {"eat", "tea", "ate"},
			},
		},
	} {
		answer := groupAnagrams(v.input)

		for i := range answer {
			result := v.result[len(answer[i])]
			for j := range answer[i] {
				if answer[i][j] != result[j] {
					t.Errorf("error: %s != %s", answer[i][j], result[j])
				}
			}
		}
	}
}
