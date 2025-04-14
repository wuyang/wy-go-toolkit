func reverse(x int) int {
	// signed 32-bit integer range is [-2^31, 2^31 - 1]
	// math.MaxInt32 = 2,147,483,647
	// math.MinInt32 = -2,147,483,648
	result := 0
	for x != 0 {
		digit := x % 10
		x /= 10
		if (result > math.MaxInt32/10 ||
			(result == math.MaxInt32/10 && digit > 7)) ||
			(result < math.MinInt32/10 ||
				(result == math.MinInt32/10 && digit < -8)) {
			return 0
		}
		result = result*10 + digit
	}
	return result
}
