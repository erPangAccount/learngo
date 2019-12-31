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
	node.Right.Traverse()
}

func (node *Node) TraceFunc(f func(*Node)) {
	if node == nil {
		return
	}

	node.Left.TraceFunc(f)
	f(node)
	node.Right.TraceFunc(f)
}

func (node *Node) TraverseWithChannel() chan *Node {
	out := make(chan *Node)
	go func() {
		node.TraceFunc(func(node *Node) {
			out <- node
		})
		close(out)
	}()
	return out
}
