func lengthOfLongestSubstring(s string) int {
        char_index_map := make(map[rune]int);  // value -> last see index
        max_length , left := 0, -1 // left is the index of the elment before the substring (exclusive)
        for right, current_char := range s  {
            if index, found :=char_index_map[current_char]; found && index > left {
                left = char_index_map[current_char]
            }
            char_index_map[current_char] = right;
            window_length := right - left
            if window_length > max_length {
                max_length = window_length
            } 
        }
        return max_length
}