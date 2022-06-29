/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
func nextPermutation(nums []int) {
	i := len(nums)
	if i <= 1 {
		return
	}
	//如果长度小于1，直接返回
	a, b, c := i-2, i-1, i-1
	for a >= 0 && nums[a] >= nums[b] {
		a--
		b--
	}
	if a >= 0 {
		for nums[a] >= nums[c] {
			c--
		}
		nums[a], nums[c] = nums[c], nums[a]
	}
	for a, b := b, i-1; a < b; a, b = a+1, b-1 {
		nums[a], nums[b] = nums[b], nums[a]
	}
}

// @lc code=end

