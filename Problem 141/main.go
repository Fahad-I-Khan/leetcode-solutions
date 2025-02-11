package main  // Linked List Cycle

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
    Val  int
    Next *ListNode
}

// Function to detect if there is a cycle in the linked list
func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }

    slow := head
    fast := head

    for fast != nil && fast.Next != nil {
        slow = slow.Next        // Move slow pointer by 1 step
        fast = fast.Next.Next  // Move fast pointer by 2 steps

        if slow == fast {
            return true  // Cycle detected
        }
    }

    return false  // No cycle
}

// Helper function to create a linked list from an array and insert a cycle
func createList(arr []int, pos int) *ListNode {
    if len(arr) == 0 {
        return nil
    }

    var head, tail, cycleStart *ListNode
    nodes := make([]*ListNode, len(arr))

    for i, val := range arr {
        node := &ListNode{Val: val}
        if head == nil {
            head = node
        } else {
            tail.Next = node
        }
        tail = node
        nodes[i] = node
    }

    if pos >= 0 && pos < len(arr) {
        tail.Next = nodes[pos] // Create a cycle
    }

    return head
}

// Helper function to print the linked list (for debugging)
func printList(head *ListNode) {
    seen := map[*ListNode]bool{}
    for head != nil {
        if seen[head] {
            fmt.Print("(cycle detected) ")
            break
        }
        fmt.Print(head.Val, " -> ")
        seen[head] = true
        head = head.Next
    }
    fmt.Println("nil")
}

func main() {
    // Example 1: [3, 2, 0, -4], pos = 1 (cycle starts at node with value 2)
    list1 := createList([]int{3, 2, 0, -4}, 1)
    fmt.Println("Has cycle (example 1):", hasCycle(list1)) // Should print: true

    // Example 2: [1, 2], pos = 0 (cycle starts at node with value 1)
    list2 := createList([]int{1, 2}, 0)
    fmt.Println("Has cycle (example 2):", hasCycle(list2)) // Should print: true
}

