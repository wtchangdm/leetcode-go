package leetcode

// https://leetcode.com/problems/min-stack/
type MinStack struct {
	min []int
	seq []int
}

func (m *MinStack) Push(val int) {
	var lastMinimum int
	if len(m.seq) == 0 {
		lastMinimum = val
	} else {
		lastMinimum = m.min[len(m.min)-1]
	}

	if val < lastMinimum {
		lastMinimum = val
	}

	m.min = append(m.min, lastMinimum)
	m.seq = append(m.seq, val)
}

func (m *MinStack) Pop() {
	m.seq = m.seq[:len(m.seq)-1]
	m.min = m.min[:len(m.min)-1]
}

func (m *MinStack) Top() int {
	return m.seq[len(m.seq)-1]
}

func (m *MinStack) GetMin() int {
	return m.min[len(m.min)-1]
}
