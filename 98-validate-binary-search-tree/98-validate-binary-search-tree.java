/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    long preVaLue = Long.MIN_VALUE;   //开始设置为整数最小值
    public boolean isValidBST(TreeNode root) {
        if(root == null) {
            return true;
        }
        boolean isleftBst = isValidBST(root.left);
        if(isleftBst == false) {
            //如果左数不是搜索二叉树，那就直接返回
            return false;
        }
        if(root.val <= preVaLue) {
            //如果发现中序遍历后面的数小于preVaLue
            return false;
        }else{   //>
            preVaLue = root.val;
        }
        // boolean isrightBst = isValidBST(root.right);
        // if(isrightBst == false) {
        //     return false;
        // }else{
        //     return true;
        //  }
        return  isValidBST(root.right);
    }
}