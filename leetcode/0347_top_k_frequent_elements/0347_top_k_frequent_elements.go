package leetcode

func topKFrequent(nums []int, k int) []int {
	count := map[int]int{}
	freq := map[int][]int{}
	result := []int{}

	for _, num := range nums {
		count[num]++
	}

	for key := range count {
		times := count[key]

		if _, ok := freq[times]; !ok {
			freq[times] = []int{}
		}

		freq[times] = append(freq[times], key)
	}

	for i := len(nums); ; i-- {
		for _, n := range freq[i] {
			result = append(result, n)
			if len(result) == k {
				return result
			}
		}
	}
}
