package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
    Val  int
    Next *ListNode
}

// Function to reverse the linked list
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode = nil
    current := head

    for current != nil {
        next := current.Next // Store next node
        current.Next = prev  // Reverse the pointer
        prev = current       // Move prev to current
        current = next       // Move to the next node
    }

    return prev // Prev will be the new head
}

// Helper function to print the linked list
func printList(head *ListNode) {
    for head != nil {
        fmt.Print(head.Val, " -> ")
        head = head.Next
    }
    fmt.Println("nil")
}

func main() {
    // Create a sample linked list: 1 -> 2 -> 3 -> 4 -> nil
    head := &ListNode{Val: 1}
    head.Next = &ListNode{Val: 2}
    head.Next.Next = &ListNode{Val: 3}
    head.Next.Next.Next = &ListNode{Val: 4}

    fmt.Println("Original List:")
    printList(head)

    reversedHead := reverseList(head)
    
    fmt.Println("Reversed List:")
    printList(reversedHead)
}
