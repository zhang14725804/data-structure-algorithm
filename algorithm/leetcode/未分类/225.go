/*
	三种方法，todo：切片、标准库list和自定义list
*/

type MyStack struct {
	list []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.list = append([]int{x}, this.list...)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	res := this.list[0]
	this.list = this.list[1:]
	return res
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.list[0]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.list) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */