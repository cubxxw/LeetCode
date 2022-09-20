func shuffle(nums []int, n int) []int {
    ans := make([]int, n*2)
    for i,v := range nums[:n] {
        ans[i*2] = v
        ans[i*2 + 1] = nums[i + n]
    }
    return ans
}
