/*
 * @Description:
 * @Author: xiongxinwei 3293172751nss@gmail.com
 * @Date: 2022-10-08 18:57:07
 * @LastEditTime: 2022-10-08 19:01:23
 * @FilePath: \undefinedd:\文档\git\LeetCode\leetcode\54.螺旋矩阵.go
 * @Github_Address: https://github.com/3293172751/cs-awesome-Block_Chain
 * Copyright (c) 2022 by xiongxinwei 3293172751nss@gmail.com, All Rights Reserved. @blog: http://nsddd.top
 */
/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return []int{}
    }
    rows, columns := len(matrix), len(matrix[0])
    visited := make([][]bool, rows)
    for i := 0; i < rows; i++ {
        visited[i] = make([]bool, columns)
    }

    var (
        total = rows * columns
        order = make([]int, total)
        row, column = 0, 0
        directions = [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}
        directionIndex = 0
    )

    for i := 0; i < total; i++ {
        order[i] = matrix[row][column]
        visited[row][column] = true
        nextRow, nextColumn := row + directions[directionIndex][0], column + directions[directionIndex][1]
        if nextRow < 0 || nextRow >= rows || nextColumn < 0 || nextColumn >= columns || visited[nextRow][nextColumn] {
            directionIndex = (directionIndex + 1) % 4
        }
        row += directions[directionIndex][0]
        column += directions[directionIndex][1]
    }
    return order
}


// @lc code=end

