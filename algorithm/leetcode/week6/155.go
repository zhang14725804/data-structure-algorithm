/*
	设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
		push(x) —— 将元素 x 推入栈中。
		pop() —— 删除栈顶的元素。
		top() —— 获取栈顶元素。
		getMin() —— 检索栈中的最小元素。
*/

type MinStack struct {
	stack []item
}
/*
	😅😅😅 存储当前值和最小值
*/ 
type item struct{
	min,x int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

// 入栈
func (this *MinStack) Push(x int)  {
	// 更新最小值
	min:=x
	if len(this.stack)>0 && this.GetMin() < x {
		min = this.GetMin()
	}
	this.stack = append(this.stack,item{min:min,x:x})
}

// 出栈
func (this *MinStack) Pop()  {
	this.stack = this.stack[:len(this.stack)-1]
}

// 返回栈顶元素
func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1].x
}

// 返回栈最小值
func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1].min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */