package main

import (
	"fmt"
	tree2 "language/tree"
)

type MyTreeNode struct {
	node *tree2.Node
}

/**
扩展树的后续遍历
*/
func (myNode *MyTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := MyTreeNode{myNode.node.Left}
	left.postOrder()
	right := MyTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree2.Node

	root = tree2.Node{Value: 3}
	root.Left = &tree2.Node{}
	root.Right = &tree2.Node{5, nil, nil}
	root.Right.Left = tree2.CreateNode(2)
	fmt.Println(root) //{3 0xc000074460 0xc000074480}

	/**
	上面创建的树的图示：
		3
	0		5
		   2
	*/

	root.Traverse() //{0 <nil> <nil>} {3 0xc0000044a0 0xc0000044c0} {2 <nil> <nil>} {5 0xc0000044e0 <nil>}
	fmt.Println()

	var myRoot MyTreeNode
	myRoot.node = &root
	myRoot.postOrder() //{0 <nil> <nil>} {2 <nil> <nil>} {5 0xc000004500 <nil>} {3 0xc0000044c0 0xc0000044e0}
	fmt.Println()

	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println("Max node value:", maxNode) //Max node value: 5
}
