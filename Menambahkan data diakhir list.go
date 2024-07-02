//Menambahkan data diakhir list
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

    // Add a new node with value 6 to the end of the list
    newNode := &Node{Value: 6}
    appendNode(head, newNode)

    // Print the updated linked list
    current := head
    for current!= nil {
        fmt.Print(current.Value, " -> ")
        current = current.Next
    }
    fmt.Println()
}

func appendNode(head *Node, newNode *Node) {
    current := head
    for current.Next!= nil {
        current = current.Next
    }
    current.Next = newNode
}