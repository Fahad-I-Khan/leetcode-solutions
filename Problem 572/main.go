package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// Function to check if two trees are the same
func isSameTree(s, t *TreeNode) bool {
    if s == nil && t == nil {
        return true
    }
    if s == nil || t == nil {
        return false
    }
    return s.Val == t.Val && isSameTree(s.Left, t.Left) && isSameTree(s.Right, t.Right)
}

// Function to check if a tree is a subtree of another tree
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil {
        return false
    }
    if isSameTree(root, subRoot) {
        return true
    }
    return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

// Helper function to create a tree from a slice (for testing purposes)
func createTree(values []int) *TreeNode {
    if len(values) == 0 || values[0] == -1 {
        return nil
    }
    root := &TreeNode{Val: values[0]}
    queue := []*TreeNode{root}
    i := 1
    for len(queue) > 0 && i < len(values) {
        node := queue[0]
        queue = queue[1:]
        if i < len(values) && values[i] != -1 {
            node.Left = &TreeNode{Val: values[i]}
            queue = append(queue, node.Left)
        }
        i++
        if i < len(values) && values[i] != -1 {
            node.Right = &TreeNode{Val: values[i]}
            queue = append(queue, node.Right)
        }
        i++
    }
    return root
}

func main() {
    // Example 1
    root1 := createTree([]int{3, 4, 5, 1, 2})
    subRoot1 := createTree([]int{4, 1, 2})
    fmt.Println(isSubtree(root1, subRoot1)) // Output: true

    // Example 2
    root2 := createTree([]int{3, 4, 5, 1, 2, -1, -1, -1, -1, 0})
    subRoot2 := createTree([]int{4, 1, 2})
    fmt.Println(isSubtree(root2, subRoot2)) // Output: false
}
