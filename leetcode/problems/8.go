import (
	"math"
)

func myAtoi(s string) int {
	sign := 1
	result := 0
	i := 0
	str := []rune(s)
	for i < len(str) && str[i] == ' ' {
		i += 1
	}
	if i < len(str) && str[i] == '+' {
		i += 1
	} else if i < len(str) && str[i] == '-' {
		i += 1
		sign = -1
	}
	for i < len(str) && unicode.IsDigit(str[i]) {
		// signed 32-bit integer range is [-2^31, 2^31 - 1]
		// math.MaxInt32 = 2,147,483,647
		// math.MinInt32 = -2,147,483,648
        digit := int(str[i] - '0')
                        if (sign > 0 &&
                        (result > math.MaxInt32 / 10 ||
                                result == math.MaxInt32 / 10 &&
                                        digit >= 7)) {
                    return math.MaxInt32;
                } else if (sign < 0 &&
                        (result > -1 * (math.MinInt32 / 10) ||
                                result == -1 * (math.MinInt32 /
                                        10) &&
                                        digit >= 8)) {
                    return math.MinInt32;
                }

		result = result * 10 + digit
		i += 1
	}
    return sign * result
}
