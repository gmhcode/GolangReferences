package main

import "fmt"

type Node interface {
	SetValue(v int)
	GetValue() int
}

//SLLNode - node class
type SLLNode struct {
	next         *SLLNode
	value        int
	SNodeMessage string
}

//SetValue - set value for node
func (sNode *SLLNode) SetValue(v int) {
	sNode.value = v
}

//GetValue - get value for node
func (sNode *SLLNode) GetValue() int {
	return sNode.value
}

//NewSLLNode - create new node
func NewSLLNode() *SLLNode {
	return &SLLNode{SNodeMessage: "This is a message from the normal Node"}
}

type PowerNode struct {
	next         *PowerNode
	value        int
	PNodeMessage string
}

func (sNode *PowerNode) SetValue(v int) {
	sNode.value = v * 10
}

func (sNode *PowerNode) GetValue() int {
	return sNode.value
}

func NewPowerNode() *PowerNode {
	return &PowerNode{PNodeMessage: "This is a message from the power Node"}
}

// func main() {
// var node Node

// node = NewSLLNode()
// node.SetValue(4)
// fmt.Println("Node is of value: ", node.GetValue())

// node = NewPowerNode()
// node.SetValue(5)
// fmt.Println("Node is of value ", node.GetValue())
// //Checks if PowerNode is implementing the Node interface
// if n, ok := node.(*PowerNode); ok {
// 	fmt.Println("This is a power node of value ", n.value)
// }

// 	list := newSingleLinkedList()
// 	list.Add(3)
// 	list.Add(4)
// 	list.Add(5)
// 	list.Add(6)
// 	fmt.Println("Hello, playground", list)
// }

func main() {
	n := createNode(5)
	//typeSwitch
	switch concreten := n.(type) {
	case *SLLNode:
		fmt.Println("Type is SLLNode, message: ", concreten.SNodeMessage)
	case *PowerNode:
		fmt.Println("Type is PowerNode, message: ", concreten.PNodeMessage)
	}

	sNode := &SLLNode{}
	n = sNode
}

// func zero(xPtr *int) {
// 	*xPtr = 0
// }
// func main() {
// 	x := 5
// 	zero(&x)
// 	fmt.Println(x) // x is 0
// }

func createNode(v int) Node {
	sn := NewSLLNode()
	sn.SetValue(v)
	return sn
}
