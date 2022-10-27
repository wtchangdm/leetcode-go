package leetcode

import (
	"testing"
)

func TestWordDictionary(t *testing.T) {
	wd := Constructor()

	wd.AddWord("bad")
	wd.AddWord("dad")
	wd.AddWord("mad")

	for _, v := range []struct {
		word   string
		expect bool
	}{
		{word: "pad", expect: false},
		{word: ".ad", expect: true},
		{word: "b..", expect: true},
	} {
		if wd.Search(v.word) != v.expect {
			t.Errorf("Test failed: search %s should be %v", v.word, v.expect)
		}
	}
}
