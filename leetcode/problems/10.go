func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	} else if len(p) >= 2 && p[1] == '*' { // process *
		if len(s) == 0 {
			// drop ".*" or "a*" in p
			return isMatch(s, p[2:])
		} else if s[0] == p[0] || p[0] == '.' {
			// match none and drop "a*" or ".*" in p
			// match one char and keep "a*" or ".*" in p
			return isMatch(s, p[2:]) || isMatch(s[1:], p)
		} else {
			// no match char and drop "a*" in p
			return isMatch(s, p[2:])
		}
	} else if len(s) > 0 && (s[0] == p[0] || p[0] == '.') {
		// match one char without *
		return isMatch(s[1:], p[1:])
	} else {
		return false
	}
}
