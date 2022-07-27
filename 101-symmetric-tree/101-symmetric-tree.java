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
    public boolean isSymmetric(TreeNode root) {
        if(root == null) {
            return false;
        }
        return check(root.left,root.right);
    }
    public boolean check(TreeNode l, TreeNode r) {
        if(l == null && r == null) {
            return true;
        }
        if(l == null || r == null) {
            return l == r ? true:false;
        }
        if(l.val != r.val) {
            return false;
        }

        return check(l.left,r.right)&&check(l.right,r.left);
    } 
}