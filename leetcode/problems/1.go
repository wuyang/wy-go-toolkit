func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i, v := range nums {
        compliment := target - v;
        if compliment_index, exist := m[compliment]; exist {
            return []int{compliment_index, i}
        }
        m[v] = i
    }
    return nil
}