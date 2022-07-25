package leetcode

import "testing"

func TestTrie(t *testing.T) {
	trie := Constructor()

	trie.Insert("apple")

	word := "apple"

	if !trie.Search(word) {
		t.Errorf("%s should be able to find", word)
	}

	word = "app"

	if trie.Search(word) {
		t.Errorf("%s shouldn't exist", word)
	}

	if !trie.StartsWith(word) {
		t.Errorf("%s should start with %s", word, word)
	}

	trie.Insert(word)
	if !trie.Search(word) {
		t.Errorf("%s should be able to find", word)
	}
}
