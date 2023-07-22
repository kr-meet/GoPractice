func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }

    left := root.Left
    right := root.Right

    root.Left = right
    root.Right = left

    invertTree(root.Left)
    invertTree(root.Right)

    return root
}