package leetcode

import "fmt"

var delimiter rune = 257
var emptyString rune = 258

type Codec struct {
}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	s := ""
	for i := range strs {
		if len(strs[i]) == 0 {
			s += string(emptyString)
		} else {
			s += strs[i]
		}

		if i < len(strs)-1 {
			s += string(delimiter)
		}
	}
	return s
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	fmt.Println(strs)
	result := []string{}

	tmp := []rune{}

	for i, r := range strs {
		if r != delimiter && r != emptyString {
			tmp = append(tmp, r)
		}

		if r == delimiter && len(tmp) > 0 {
			result = append(result, string(tmp))
			tmp = []rune{}
			continue
		}

		if r == emptyString {
			result = append(result, "")
			continue
		}

		if i == len(strs)-1 {
			result = append(result, string(tmp))
			tmp = []rune{}
		}
	}

	return result
}

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));
