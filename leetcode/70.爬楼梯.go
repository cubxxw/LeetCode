/*
 * @Description:  爬楼梯
 * @Author: xiongxinwei 3293172751nss@gmail.com
 * @Date: 2022-09-30 13:02:49
 * @LastEditTime: 2022-09-30 13:16:39
 * @FilePath: \leetcode\70.爬楼梯.go
 * @Github_Address: https://github.com/3293172751/cs-awesome-Block_Chain
 * Copyright (c) 2022 by xiongxinwei 3293172751nss@gmail.com, All Rights Reserved. @blog: http://nsddd.top
 */
/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 	// 爬楼梯
 if n = 4  ==> 1 + （1 + 1 * 2 + 1 * 3 + ... + n）
	1 + 1 + 1 + 1
	1 + 1 + 2 * 3
	1 + 3 * 2 
	4 
	*/

// @lc code=start
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	sqrt5 := math.Sqrt(5)
    pow1 := math.Pow((1+sqrt5)/2, float64(n+1))
    pow2 := math.Pow((1-sqrt5)/2, float64(n+1))
    return int(math.Round((pow1 - pow2) / sqrt5))
}
// @lc code=end

