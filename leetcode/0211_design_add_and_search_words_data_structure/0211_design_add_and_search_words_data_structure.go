package leetcode

type WordDictionary struct {
	dict   map[rune]*WordDictionary
	isWord bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		dict: map[rune]*WordDictionary{},
	}
}

func (this *WordDictionary) AddWord(word string) {
	curr := this

	for _, char := range word {
		if _, found := curr.dict[char]; !found {
			newNode := Constructor()
			curr.dict[char] = &newNode
		}
		curr = curr.dict[char]
	}

	curr.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
	curr := this

	for i := 0; i < len(word); i++ {
		char := rune(word[i])

		if char == '.' {
			for _, subDict := range curr.dict {
				found := subDict.Search(word[i+1:])
				if found {
					return true
				}
			}
			return false
		} else {
			if _, found := curr.dict[char]; !found {
				return false
			}

			curr = curr.dict[char]
		}
	}

	return curr.isWord
}
