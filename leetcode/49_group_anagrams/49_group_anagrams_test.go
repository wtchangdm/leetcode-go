package leetcode

import "testing"

func TestGroupAnagrams(t *testing.T) {

	for _, v := range []struct {
		input  []string
		result [][]string
	}{
		{
			input:  []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			result: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
	} {
		answer := groupAnagrams(v.input)

		for i := range answer {
			for j := range answer[i] {
				if answer[i][j] != v.result[i][j] {
					t.Errorf("error: %s != %s", answer[i][j], v.result[i][j])
				}
			}
		}
	}
}
