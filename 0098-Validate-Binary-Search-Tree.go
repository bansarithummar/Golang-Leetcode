/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    return validate(root, nil, nil)
}

func validate(node *TreeNode, lower *int, upper *int) bool {
	if node == nil {
		return true
	}
	
	val := node.Val
	
	if lower != nil && val <= *lower {
		return false
	}
	if upper != nil && val >= *upper {
		return false
	}
	
	if !validate(node.Right, &val, upper) {
		return false
	}
	if !validate(node.Left, lower, &val) {
		return false
	}
	
	return true
}
