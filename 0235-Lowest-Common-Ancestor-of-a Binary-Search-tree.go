235. Lowest Common Ancestor of a Binary Search tree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    for root != nil {
        if p.Val > root.Val && q.Val > root.Val {
            root = root.Right
        } else if p.Val < root.Val && q.Val < root.Val {
            root = root.Left
        } else {
            return root
        }
    }
    return nil
}
