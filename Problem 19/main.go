package main // Leetcode 19 

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
    Val  int
    Next *ListNode
}

// Function to remove the n-th node from the end of the list
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    // Create a dummy node to simplify edge cases
    dummy := &ListNode{Next: head}
    fast := dummy
    slow := dummy

    // Move fast pointer n+1 steps ahead
    for i := 0; i <= n; i++ {
        fast = fast.Next
    }

    // Move both fast and slow pointers until fast reaches the end
    for fast != nil {
        fast = fast.Next
        slow = slow.Next
    }

    // Remove the n-th node from end
    slow.Next = slow.Next.Next

    // Return the new head, which is the next node of dummy
    return dummy.Next
}

// Helper function to print the linked list
func printList(head *ListNode) {
    for head != nil {
        fmt.Print(head.Val, " -> ")
        head = head.Next
    }
    fmt.Println("nil")
}

// Helper function to create a linked list from an array
func createList(arr []int) *ListNode {
    if len(arr) == 0 {
        return nil
    }
    head := &ListNode{Val: arr[0]}
    current := head
    for _, val := range arr[1:] {
        current.Next = &ListNode{Val: val}
        current = current.Next
    }
    return head
}

func main() {
    // Example: head = [1,2,3,4,5], n = 2
    head := createList([]int{1, 2, 3, 4, 5})
    n := 2
    newHead := removeNthFromEnd(head, n)
    fmt.Print("Modified List: ")
    printList(newHead)
}
