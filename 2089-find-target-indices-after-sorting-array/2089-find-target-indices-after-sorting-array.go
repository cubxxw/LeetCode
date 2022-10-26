/*
 * @Description: 二分查找
 * @Author: xiongxinwei 3293172751nss@gmail.com
 * @Date: 2022-10-26 20:46:14
 * @LastEditTime: 2022-10-26 20:54:26
 * @FilePath: \coded:\文档\git\LeetCode\leetcode\2089.找出数组排序后的目标下标.go
 * @Github_Address: https://github.com/3293172751/cs-awesome-Block_Chain
 * Copyright (c) 2022 by xiongxinwei 3293172751nss@gmail.com, All Rights Reserved. @blog: http://nsddd.top
 */
/*
 * @lc app=leetcode.cn id=2089 lang=golang
 *
 * [2089] 找出数组排序后的目标下标
 go sort排序
*/

// @lc code=start
func targetIndices(nums []int, target int) []int {
	less, equal := 0, 0
	for _, num := range nums {
		if num < target {
			less++
		} else if num == target {
			equal++
		}
	}
	ans := make([]int, equal)
	for i := range ans {
		ans[i] = less + i
	}
	return ans
}

// @lc code=end

