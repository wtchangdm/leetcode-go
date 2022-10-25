package leetcode

import "sort"

func canAttendMeetings(intervals [][]int) bool {
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i < len(intervals)-1; i++ {
		currentEnd := intervals[i][len(intervals[i])-1]
		nextStart := intervals[i+1][0]

		if nextStart < currentEnd {
			return false
		}
	}

	return true
}
