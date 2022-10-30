package leetcode

import (
	"testing"
)

func TestNumIslands(t *testing.T) {
	for _, v := range []struct {
		input  [][]byte
		expect int
	}{
		{
			input: [][]byte{
				{byte('1'), byte('1'), byte('1'), byte('1'), byte('0')},
				{byte('1'), byte('1'), byte('0'), byte('1'), byte('0')},
				{byte('1'), byte('1'), byte('0'), byte('0'), byte('0')},
				{byte('0'), byte('0'), byte('0'), byte('0'), byte('0')},
			},
			expect: 1,
		},
	} {
		result := numIslands(v.input)
		if result != v.expect {
			t.Errorf("result %v != %v", result, v.expect)
		}
	}
}
