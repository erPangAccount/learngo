package main

import (
	"fmt"
	"tree"
)

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = tree.CreateNode(2)
	fmt.Println(root) //{3 0xc000074460 0xc000074480}

	/**
	上面创建的树的图示：
		3
	0		5
		   2
	*/
	root.Traverse() //{0 <nil> <nil>} {3 0xc0000044a0 0xc0000044c0} {2 <nil> <nil>} {5 0xc0000044e0 <nil>}
	fmt.Println()

	nodeCount := 0
	root.TraceFunc(func(node *tree.Node) { //{0 <nil> <nil>} {3 0xc0000044c0 0xc0000044e0} {0 <nil> <nil>}
		node.Print()
		nodeCount++
	})
	fmt.Println("nodeCount: ", nodeCount) //nodeCount:  3
}
