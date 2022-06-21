package leetcode

func reverseBits(num uint32) uint32 {
	var result uint32 = 0

	iterations := 32

	for i := 0; i < iterations; i++ {
		result = result << 1
		var bit uint32 = 0

		if num%2 == 1 {
			bit = 1
		}

		result += bit

		num = num >> 1

	}

	return result
}
