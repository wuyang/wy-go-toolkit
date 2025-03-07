func longestPalindrome(s string) string {
  start := 0 // start index of the longest palindrome so far
  length := 0 // length of the longest palindrome so far
  for i := 0; i < len(s); i++ {
    // expand odd palindrome like aba from center i
    oddLength := expandFromCenter(s, i, i)
    // expand even palindrome like abba from center i and i+1
    evenLength := expandFromCenter(s, i, i + 1)
    maxLength := 0
    if oddLength > evenLength {
      maxLength = oddLength
    } else {
      maxLength = evenLength
    }
    if maxLength > length {
      length = maxLength
      start = i - (maxLength - 1) / 2
    }
  }
  return s[start:start + length]
}

func expandFromCenter(s string, left int, right int) int {
  for left >= 0 && right < len(s) && s[left] == s[right] {
    left--
    right++
  }
  // [left + 1, right - 1] is the longest palindrome
  return right - 1 - (left + 1) + 1
}