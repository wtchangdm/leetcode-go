package leetcode

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxArea(height []int) int {
	maxArea := 0
	left, right := 0, len(height)-1

	for left < right {
		length := right - left
		minHeight := min(height[left], height[right])
		area := length * minHeight

		if area > maxArea {
			maxArea = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}
