/* 简单 */
type CustomStack struct {
	Stack []int
	Length int
	MaxSize int
}


func Constructor(maxSize int) CustomStack {
	return CustomStack{
		Stack:   nil,
		Length:  0,
		MaxSize: maxSize,
	}
}


func (this *CustomStack) Push(x int)  {
	if this.Length == this.MaxSize{
		return
	}
	this.Length++
	this.Stack = append(this.Stack,x)
}


func (this *CustomStack) Pop() int {
	if this.Length == 0 {
		return -1
	}
	this.Length--
	x:=this.Stack[len(this.Stack)-1]
	this.Stack =  this.Stack[:len(this.Stack)-1]
	return x
}


func (this *CustomStack) Increment(k int, val int)  {
	for i:=0;i<k && i<this.Length ;i++  {
		this.Stack[i]+=val
	}
}