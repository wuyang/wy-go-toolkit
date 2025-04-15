func maxArea(height []int) int {
	maxArea := 0
	left := 0
	right := len(height) - 1
	for right > left {
		maxArea = max(maxArea, (right-left)*min(height[left], height[right]))
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return maxArea
}
