package leetcode

import "testing"

func TestCodeC(t *testing.T) {
	var codec Codec
	codec.Decode(codec.Encode([]string{""}))

	for _, v := range []struct {
		input  []string
		result []string
	}{
		{input: []string{"I", "am", "test"}, result: []string{"I", "am", "test"}},
		{input: []string{"", ""}, result: []string{"", ""}},
	} {
		answer := codec.Decode(codec.Encode(v.input))
		for i := range answer {
			if answer[i] != v.result[i] {
				t.Errorf("error: %s != %s", answer[i], v.result[i])
			}
		}
	}
}
