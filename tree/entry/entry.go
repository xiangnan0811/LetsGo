package main

import (
	"fmt"

	"imooc.com/xiangnan/learngo/tree"

	"golang.org/x/tools/container/intsets"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func testSparse() {
	s := intsets.Sparse{}
	s.Insert(1)
	s.Insert(1000)
	s.Insert(1000000)
	fmt.Println(s.Has(1000))
	fmt.Println(s.Has(100000))
}

func main() {
	// 结构的创建
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse() // 遍历
	fmt.Println()

	// 值接收者 vs 指针接收者
	/* 1. 要改变内容必须使用指针接收者
	   2. 结构过大也考虑使用指针接收者
	   3. 一致性：如有指针接收者，最好都是指针接收者
	*/
	myNode := myTreeNode{&root}
	myNode.postOrder()
	fmt.Println()

	testSparse()

	// Go语言闭包的应用：为函数实现接口
	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node count:", nodeCount)
}
