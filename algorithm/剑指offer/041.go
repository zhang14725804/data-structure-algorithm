/*
	包含min函数的栈（单调栈）
	stk存放所有数据，minStk存放最小值
*/
type MinStack struct {
    stk,minStk []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(x int)  {
	this.stk = append(this.stk,x)
	if len(this.minStk) == 0 || this.minStk[len(this.minStk)-1] >= x {
		this.minStk = append(this.minStk,x)
	}
}


func (this *MinStack) Pop()  {
    if this.Top() == this.minStk[len(this.minStk)-1]{
		this.minStk = this.minStk[:len(this.minStk)-1]
	}
	this.stk = this.stk[:len(this.stk)-1]
}


func (this *MinStack) Top() int {
    return this.stk[len(this.stk)-1]
}


func (this *MinStack) GetMin() int {
    return this.minStk[len(this.minStk)-1]
}