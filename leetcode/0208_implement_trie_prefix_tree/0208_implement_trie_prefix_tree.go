package leetcode

type TrieNode struct {
	children    map[rune]*TrieNode
	isEndOfWord bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
		},
	}
}

func (this *Trie) Insert(word string) {
	current := this.root

	for _, r := range word {
		if _, found := current.children[r]; !found {
			current.children[r] = &TrieNode{
				children: make(map[rune]*TrieNode),
			}
		}

		current = current.children[r]
	}
	current.isEndOfWord = true
}

func (this *Trie) Search(word string) bool {
	current := this.root

	for _, r := range word {
		if _, found := current.children[r]; !found {
			return false
		}

		current = current.children[r]
	}
	return current.isEndOfWord
}

func (this *Trie) StartsWith(prefix string) bool {
	current := this.root

	for _, r := range prefix {
		if _, found := current.children[r]; !found {
			return false
		}

		current = current.children[r]
	}
	return true
}
