 

import (
	"code-data-structure/Single_link"
	"code-data-structure/double_link"
	"fmt"
)

/*
	build command-line-arguments: cannot find module for path _/D_/aprivate/code-data-structure/Single_link
	参考：https://blog.csdn.net/sinat_23156865/article/details/100655475
	https://blog.csdn.net/nia305/article/details/89966314

	链表：适合频繁插入删除
	数组：适用于查询
	window进程管理：双链表

	求链表位置：中间节点，三分之一位置，四分之一位置（快慢指针）
*/
func main1() {
	list := Single_link.NewSingleLinkList()
	node1 := Single_link.NewSingleLinkNode(1)
	node2 := Single_link.NewSingleLinkNode(2)
	node3 := Single_link.NewSingleLinkNode(3)
	node4 := Single_link.NewSingleLinkNode(4)
	node5 := Single_link.NewSingleLinkNode(5)
	//list.InsertFront(node1)
	//list.InsertFront(node2)
	//list.InsertFront(node3)
	list.InsertBack(node1)
	list.InsertBack(node2)
	list.InsertBack(node3)
	list.InsertBefore(1, node4)
	list.InsertAfter(3, node5)
	fmt.Println(list)
	fmt.Println(list.GetNodeAtIndex(2))
	list.DeleteNode(2)
	fmt.Println(list)
	list.DeleteNodeAtIndex(1)
	fmt.Println(list)
	fmt.Println(list.GetMiddle())
	list.Reverse()
	fmt.Println(list)
}

/*
	双链表
*/
func main2() {
	list := double_link.NewDoubleLink()
	node1 := double_link.NewDoublelinkNode(1)
	node2 := double_link.NewDoublelinkNode(2)
	node3 := double_link.NewDoublelinkNode(3)
	node4 := double_link.NewDoublelinkNode(4)
	list.InsertBefore(node1)
	list.InsertBefore(node2)
	list.InsertAtfer(node3)
	list.InsertAtfer(node4)
	fmt.Println(list.String())
}

func main() {
	list := double_link.NewDoubleLink()
	node1 := double_link.NewDoublelinkNode(1)
	node2 := double_link.NewDoublelinkNode(2)
	node3 := double_link.NewDoublelinkNode(3)
	node4 := double_link.NewDoublelinkNode(4)
	list.InsertBefore(node1)
	list.InsertBefore(node2)
	list.InsertBefore(node3)
	list.InsertBefore(node4)
	fmt.Println(list.String())
	node5 := double_link.NewDoublelinkNode(5)
	node6 := double_link.NewDoublelinkNode(5)
	list.InsertNodeBefore(node3, node5)
	list.InsertNodeAfter(node3, node6)
	fmt.Println(list.String())
}