/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    var res []int 
    var x  func(node *TreeNode)
    x = func(node *TreeNode){
        if node == nil {
            return 
        }
        // 中序遍历
        x(node.Left)  //前序遍历
        res = append(res,node.Val)   
        x(node.Right)
    }
    x(root)
    return res
}