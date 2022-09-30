/*
 * @Description: 
 * @Author: xiongxinwei 3293172751nss@gmail.com
 * @Date: 2022-09-30 22:34:24
 * @LastEditTime: 2022-09-30 23:12:36
 * @FilePath: \HTMLd:\文档\git\LeetCode\leetcode\22.括号生成.go
 * @Github_Address: https://github.com/3293172751/cs-awesome-Block_Chain
 * Copyright (c) 2022 by xiongxinwei 3293172751nss@gmail.com, All Rights Reserved. @blog: http://nsddd.top
 */
/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	var ans []string
	var dfs func(left, right int, s string)
	dfs = func(left, right int, s string) {
		if left == n && right == n {
			ans = append(ans, s)
			return
		}
		
		if left < n {
			dfs(left+1, right, s+"(")
		}

		if right < left {
			dfs(left, right+1, s+")")
		}
	}
	dfs(0, 0, "")
	return ans
}
// @lc code=end

