/*
	84. Largest Rectangle in Histogram
	todos：：单调栈问题（不懂）

	最值问题的核心：：如何枚举所有情况

	找出左右边离他最近的矩形，在求最大值
*/
func largestRectangleArea(heights []int) int {
	n:=len(heights)
	left:=make([]int,n)
	right:=make([]int,n)

	stk := &stack{}
	// 左边
	for i:=0;i<n;i++{
		for stk.size() >0 && heights[stk.top()] >= heights[i]{
			stk.pop()
		}
		if stk.size == 0 {
			left[i] = -1
		}else{
			left[i] = stk.top()
		}
		stk.push(i)
	}
	// 清空栈
	stk = &stack{}
	// 右边
	for i:=n-1;i>=0;i--{
		for stk.size() >0 && heights[stk.top()] >= heights[i]{
			stk.pop()
		}
		if stk.size == 0 {
			right[i] = n
		}else{
			right[i] = stk.top()
		}
		stk.push(i)
	}

	res:=0
	for i:=0;i<n;i++{
		res = max(res,heights[i]*(right[i]-left[i]-1))
	}
	return res
}

// 栈
type stack struct{
	x []int
}
// 入栈
func (this *stack) push(x int){
	this.x = append(this.x,x)
}

// 入栈
func (this *stack) size() int{
	return len(this.x)
}

// 出栈
func (this *stack) pop(){
	this.x= this.x[:len(this.x)-1]
}
// 返回栈顶元素
func (this *stack) top() int{
	return this.x[len(this.x)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}