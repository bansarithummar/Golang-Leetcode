/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    _, balanced := checkBalance(root)
    return balanced
}

func checkBalance(node *TreeNode) (int, bool) {
    if node == nil {
        return 0, true
    }
    
    leftDepth, leftBalanced := checkBalance(node.Left)
    rightDepth, rightBalanced := checkBalance(node.Right)
    
    if !leftBalanced || !rightBalanced || math.Abs(float64(leftDepth-rightDepth)) > 1 {
        return 0, false
    }
    
    return max(leftDepth, rightDepth) + 1, true
}
