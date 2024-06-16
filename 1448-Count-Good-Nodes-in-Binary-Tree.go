/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    return dfs(root, root.Val)
}

func dfs(node *TreeNode, maxVal int) int {
	if node == nil {
		return 0
	}
	good := 0
	if node.Val >= maxVal {
		good = 1
	}
	newMaxVal := max(maxVal, node.Val)
	return good + dfs(node.Left, newMaxVal) + dfs(node.Right, newMaxVal) 
}
