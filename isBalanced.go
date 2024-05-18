110. Balanced Binary Tree

func isBalanced(root *TreeNode) bool {
    _, balanced := checkBalance(root)
    return balanced
}

// checkBalance returns the height of the node and whether it's balanced.
func checkBalance(node *TreeNode) (int, bool) {
    if node == nil {
        return 0, true
    }

    leftHeight, leftBalanced := checkBalance(node.Left)
    rightHeight, rightBalanced := checkBalance(node.Right)

    // Check if the current node is balanced and if both subtrees are balanced
    height := 1 + max(leftHeight, rightHeight)
    isBalanced := leftBalanced && rightBalanced && math.Abs(float64(leftHeight-rightHeight)) <= 1

    return height, isBalanced    
}
