package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

/**
方法和普通函数没什么区别，只是语法不同，都是值传递
func (接收者) 函数名 {
	函数体
}
*/
func (node Node) Print() {
	fmt.Print(node, " ")
}

func (node Node) SetValue(value int) {
	node.Value = value
}

func (node *Node) PtrSetValue(value int) {
	if node == nil {
		fmt.Println("Node is nil")
		return
	}

	node.Value = value
}

/**
工厂函数
*/
func CreateNode(value int) *Node {
	return &Node{Value: value}
}
