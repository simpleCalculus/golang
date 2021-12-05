func getConcatenation(nums []int) []int {
    nums2 := make([]int, 2*len(nums))
    for i, _ := range nums {
        nums2[i], nums2[i + len(nums)] = nums[i], nums[i]
    }
    return nums2
}
