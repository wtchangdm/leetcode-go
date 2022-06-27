package leetcode

import "testing"

func TestCodeC(t *testing.T) {
	var codec Codec
	codec.Decode(codec.Encode([]string{""}))

	for _, v := range []struct {
		input  []string
		answer []string
	}{
		{input: []string{"I", "am", "test"}, answer: []string{"I", "am", "test"}},
		{input: []string{"", ""}, answer: []string{"", ""}},
	} {
		result := codec.Decode(codec.Encode(v.input))
		for i := range result {
			if result[i] != v.answer[i] {
				t.Errorf("error: %s != %s", result[i], v.answer[i])
			}
		}
	}
}
