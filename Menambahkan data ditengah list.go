Menambahkan data ditengah list
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

    // Add a new node with value 10 at index 2
    newNode := &Node{Value: 10}
    insertNode(head, 2, newNode)

    // Print the updated linked list
    current := head
    for current != nil {
        fmt.Print(current.Value, " -> ")
        current = current.Next
    }
    fmt.Println()
}

func insertNode(head *Node, index int, newNode *Node) {
    current := head
    for i := 0; i < index-1; i++ {
        current = current.Next
    }
    newNode.Next = current.Next
    current.Next = newNode
}