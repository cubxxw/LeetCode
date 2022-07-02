/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	//快慢指针方法
	show := 0 //慢指针
	for fort := 0; fort < len(nums); fort++ {
		//外层循环整个数组，判断是否找到目标的数组
		if nums[fort] != val {
			//当不是目标数组的时候，我们就将其数字记录
			nums[show] = nums[fort]
			show++
		}
	}
	return show
}

// @lc code=end

