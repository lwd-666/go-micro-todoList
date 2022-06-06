package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func main() {
	head := &Node{}
	var i int
	fmt.Scanln(&i)
	n1 := &Node{
		Data: i,
	}
	Insert(head, n1)

}
func Insert(head, n *Node) {
	tmp := head
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = n
}
