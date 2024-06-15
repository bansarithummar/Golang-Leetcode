/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    if root == nil {
		return []int{}
	}
	var result []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		var rightmostNode int
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			rightmostNode = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, rightmostNode)
	}
	return result    
}
