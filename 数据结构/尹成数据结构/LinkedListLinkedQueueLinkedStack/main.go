 

import "fmt"

/*
	链式存储结构（没有堆栈溢出问题）
*/

func main1() {
	/*
		这样写的时候在根目录下执行go build就可以执行了
		直接运行main.go报错：：undefined: Node
	*/
	node1 := new(Node)
	node2 := new(Node)
	node3 := new(Node)

	node1.data = 1
	node1.Next = node2
	node2.data = 2
	node2.Next = node3
	node3.data = 3
	fmt.Println(node1.data)
	fmt.Println(node2.data)
	fmt.Println(node3.data)

}

/*

	todos：：
	大型项目用链式存储（重量级），玩具用顺序表就行（轻量级）
	队列是广度遍历
	栈是深度遍历

	（实现文件夹遍历，搜索文件）
	数组实现栈，队列
	链表实现栈，队列
*/
func main2() {
	// 链式栈
	stack := NewStack()
	for i := 0; i < 10; i++ {
		stack.Push(i)
	}
	for data := stack.Pop(); data != nil; data = stack.Pop() {
		fmt.Println(data)
	}
}

func main() {
	// 链式队列
	ql := NewLinkQueue()
	for i := 0; i < 10; i++ {
		ql.EnQueue(i)
	}
	fmt.Println(ql.Length())
	for data := ql.DeQueue(); data != nil; data = ql.DeQueue() {
		fmt.Println(data)
	}
}