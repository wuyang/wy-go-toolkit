func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
        n1 := len(nums1)
        n2 := len(nums2);
        // n1 needs to be smaller than n2
        // otherwise p2 could be out of nums2 in edage cases.
        if n1 > n2 {
            return findMedianSortedArrays(nums2, nums1);
        }

        left := 0;
        // If use n1 - 1, the round of -1/2 will be hard to handle 
        right := n1;
        for left <= right {
            // p1 is both the index of first right side element in nums1
            // and nums1's left side element count.
            // p2 is both the index of first right side element in nums2
            // and nums2's left side element count.
            // So, p1 + p2 = (n1 + n2 + 1 ) / 2 = total left side element count.
            // If use the index of last left side element, the edge cases will
            // much more complicated. 
            p1 := (left + right) / 2; 
            p2 := (n1 + n2 + 1 ) / 2  - p1; 
            // Value of the last left side element in nums1 
            left_max_1 := math.MinInt
            if p1 - 1 >= 0 {
                left_max_1 = nums1[p1 - 1]
             }
            // Value of the first right side element in nums1 
            right_min_1 := math.MaxInt
            if p1 < n1 {
                right_min_1 = nums1[p1]
             }
            // Value of the last left side element in nums2
            left_max_2 := math.MinInt
            if p2 - 1 >= 0 {
                left_max_2 = nums2[p2 - 1]
             } 
            // Value of the first right side element in nums2 
            right_min_2 := math.MaxInt
            if p2 < n2 {
                right_min_2 = nums2[p2]
             }
            if left_max_1 > right_min_2 {
                right = p1 - 1;
            } else if left_max_2 > right_min_1 {
                left = p1 + 1;
            } else {
                if (n1 + n2) % 2 == 1 { // odd
                    return math.Max(float64(left_max_1), float64(left_max_2));
                } else { // even
                    return (math.Max(float64(left_max_1), float64(left_max_2)) +
                      math.Min(float64(right_min_1), float64(right_min_2))) / 2;
                }
            }
        }
        return 0; // Not possible

}