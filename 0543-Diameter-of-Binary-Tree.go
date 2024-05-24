543. Diameter of Binary Tree


func diameterOfBinaryTree(root *TreeNode) int {
    maxDiameter := 0
    var depth func(node *TreeNode) int
    depth = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        leftDepth := depth(node.Left)
        rightDepth := depth(node.Right)
        // Update the maximum diameter
        maxDiameter = max(maxDiameter, leftDepth + rightDepth)
        // Return the maximum depth so it can be used by the parent call
        return max(leftDepth, rightDepth) + 1
    }
    depth(root)
    return maxDiameter    
}
