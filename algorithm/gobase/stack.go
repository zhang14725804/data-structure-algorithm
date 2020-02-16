package main

import (
	"fmt"
)

func main() {
	var left Stack
	left.push(TreeNode{1,nil,nil})
	fmt.Println(left)
	
	aa:=left.pop()
	fmt.Println(aa)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 用slice模拟栈
type Stack []TreeNode

func (s *Stack) push(node TreeNode) {
	*s = append(*s, node)
}
stac
func (s *Stack) pop() *TreeNode {
	theStack := *s
	node := &TreeNode{}
	if len(theStack) == 0 {
		return node 
	}
	node = &theStack[len(theStack)-1]
	*s = theStack[0 : len(theStack)-1]
	return node
}