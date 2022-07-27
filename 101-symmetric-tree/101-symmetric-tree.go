// https://leetcode.cn/problems/symmetric-tree/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return check(root.Left,root.Right)
    //下面判断两边的二叉树是否相等

}
func check(Left *TreeNode, Right *TreeNode) bool{
        if Left==nil || Right == nil {
            return (Left==Right)
        }
        if Left==nil && Right==nil {
            return true
        }
        if Left.Val != Right.Val {
            //如果左右子树的值不相同
            return false
        }
        return check(Left.Right, Right.Left)&&check(Left.Left,Right.Right)
    }