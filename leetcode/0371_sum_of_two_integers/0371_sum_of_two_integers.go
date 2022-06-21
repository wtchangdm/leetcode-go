package leetcode

func getSum(a int, b int) int {
	for b != 0 {
		tmp := (a & b) << 1 // take carry
		a = a ^ b
		b = tmp
	}
	return a
}
