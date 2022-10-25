package leetcode

import (
	"testing"
)

func TestCanAttendMeetings(t *testing.T) {
	for _, v := range []struct {
		question [][]int
		answer   bool
	}{
		{
			question: [][]int{{0, 30}, {5, 10}, {15, 20}},
			answer:   false,
		},
	} {
		result := canAttendMeetings(v.question)
		if result != v.answer {
			t.Error("error:", v.answer)
		}
	}
}
