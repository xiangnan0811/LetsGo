package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}

func (node Node) Print() { // 为结构定义方法
	fmt.Print(node.Value, " ")
}

func (node *Node) SetValue(value int) { // 使用指针作为方法接收者
	if node == nil {
		fmt.Println("Setting Value to nil node. Ignored")
		return
	}
	node.Value = value
}
