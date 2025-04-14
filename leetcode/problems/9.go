func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var t int = x   // temp variable to keep x unchanged
	var r int64 = 0 // reversed x
	for t > 0 {
		r = r*10 + int64(t%10)
		t /= 10
	}
	return x == int(r)
}
