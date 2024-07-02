// Menghapus data ditengah list
package main

import "fmt"

type Node struct {
    Value int
    Next  *Node
}

func main() {
    // Create a sample linked list: 1 -> 2 -> 3 -> 4 -> 5
    head := &Node{Value: 1}
    head.Next = &Node{Value: 2}
    head.Next.Next = &Node{Value: 3}
    head.Next.Next.Next = &Node{Value: 4}
    head.Next.Next.Next.Next = &Node{Value: 5}

    // Print the original linked list
    fmt.Println("Original list:")
    printList(head)

    // Remove the node with value 3 from the list
    head = removeNode(head, 3)

    // Print the updated linked list
    fmt.Println("Updated list:")
    printList(head)
}

func removeNode(head *Node, value int) *Node {
    if head == nil {
        return nil
    }

    if head.Value == value {
        return head.Next
    }

    current := head
    for current.Next != nil {
        if current.Next.Value == value {
            current.Next = current.Next.Next
            return head
        }
        current = current.Next
    }

    return head
}

func printList(head *Node) {
    for head != nil {
        fmt.Print(head.Value, " -> ")
        head = head.Next
    }
    fmt.Println()
}