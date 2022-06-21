package leetcode

// https://leetcode.com/problems/max-stack/submissions/
type MaxStack struct {
	seq []int
	max []int
}

func (m *MaxStack) Push(x int) {
	var lastMax int

	if len(m.seq) == 0 {
		lastMax = x
	} else {
		lastMax = m.max[len(m.max)-1]
	}

	if x > lastMax {
		lastMax = x
	}

	m.seq = append(m.seq, x)
	m.max = append(m.max, lastMax)
}

func (m *MaxStack) Pop() int {
	result := m.seq[len(m.seq)-1]

	m.seq = m.seq[:len(m.seq)-1]
	m.max = m.max[:len(m.max)-1]

	return result
}

func (m *MaxStack) Top() int {
	return m.seq[len(m.seq)-1]
}

func (m *MaxStack) PeekMax() int {
	return m.max[len(m.max)-1]
}

func (m *MaxStack) PopMax() int {
	max := m.PeekMax()

	tmp := []int{}

	for m.Top() != max {
		tmp = append(tmp, m.Pop())
	}

	m.Pop()

	for i := len(tmp) - 1; i >= 0; i-- {
		m.Push(tmp[i])
	}

	return max
}

/**
* Your MaxStack object will be instantiated and called as such:
* obj := Constructor();
* obj.Push(x);
* param_2 := obj.Pop();
* param_3 := obj.Top();
* param_4 := obj.PeekMax();
* param_5 := obj.PopMax();
 */

// func main() {
// 	obj := Constructor()
// 	obj.Push(1)
// 	obj.Push(2)
// 	fmt.Println(obj.Pop())
// 	fmt.Println(obj.Top())
// 	fmt.Println(obj.PeekMax())
// 	fmt.Println(obj.PopMax())
// }
