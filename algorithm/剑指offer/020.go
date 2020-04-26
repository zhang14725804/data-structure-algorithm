/*
	用两个栈实现队列
*/
type MyQueue struct {
	Stack []int
	Cache []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		Stack: nil,
		Cache: nil,
	}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
    this.Stack = append(this.Stack, x)
}

func Copy(this *MyQueue,out bool) {
	// stack -> cache
	if out{
		for len(this.Stack) > 0{
			this.Cache = append(this.Cache, this.Stack[len(this.Stack)-1])
			this.Stack = this.Stack[:len(this.Stack)-1]
		}
	}else{
		for len(this.Cache) > 0{
			this.Stack = append(this.Stack, this.Cache[len(this.Cache)-1])
			this.Cache = this.Cache[:len(this.Cache)-1]
		}
	}
    
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	Copy(this,true)
	res:=this.Cache[len(this.Cache)-1]
	this.Cache = this.Cache[:len(this.Cache)-1]
	Copy(this,false)
	return res
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
    Copy(this,true)
	res:=this.Cache[len(this.Cache)-1]
	Copy(this,false)
	return res
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    return len(this.Stack) == 0
}