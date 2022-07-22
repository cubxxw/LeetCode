/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 */

// @lc code=start
func sortedSquares(nums []int) []int {
	show := 0 //慢指针
	for i := 0; i < len(nums); i++ {
		nums[show] = nums[i] * nums[i]
		show++
	}
	sort.Ints(nums)
	return nums
}

// @lc code=end

