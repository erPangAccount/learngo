package tree

/**
树遍历 中序
*/
func (node *Node) Traverse() {
	if node == nil {
		return
	}

	node.Left.Traverse()
	node.Print()
	node.Left.Traverse()
}

func (node *Node) TraceFunc(f func(*Node)) {
	if node == nil {
		return
	}

	node.Left.TraceFunc(f)
	f(node)
	node.Left.TraceFunc(f)
}
