package leetcode

import (
	"sort"
	"strings"
)

func sortString(s string) string {
	split := strings.Split(s, "")
	sort.Strings(split)
	return strings.Join(split, "")
}

func groupAnagrams(strs []string) [][]string {
	if len(strs) <= 1 {
		return [][]string{strs}
	}

	cache := map[string][]string{}

	for _, str := range strs {
		sorted := sortString(str)
		if _, found := cache[sorted]; !found {
			cache[sorted] = []string{}
		}
		cache[sorted] = append(cache[sorted], str)
	}

	result := [][]string{}

	for key := range cache {
		result = append(result, cache[key])
	}

	return result
}
