/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */

// @lc code=start
func search(nums []int, target int) int {
	left := 0
	reght := len(nums) - 1
	for left <= reght {
		mid := (left + reght) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			reght = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

// @lc code=end

