/*
	42. Trapping Rain Water（84题的进阶版）
	给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

	思路1：按行求（会超时）
	思路2：按列求
	思路3：动态规划
	思路4：双指针
	思路5：单调栈
*/
func trap(height []int) int {
	res := 0
	stk := &stack{}

	for i := 0; i < len(height); i++ {
		last := 0
		for stk.size() > 0 && height[stk.top()] < height[i] {
			t := stk.top()
			stk.pop()
			res += (i - t - 1) * (height[t] - last)
			last = height[t]
		}
		if stk.size() > 0 {
			res += (i - stk.top() - 1) * (height[i] - last)
		}
		stk.push(i)
	}
	return res
}

// 栈
type stack struct {
	x []int
}

// 入栈
func (this *stack) push(x int) {
	this.x = append(this.x, x)
}

// 入栈
func (this *stack) size() int {
	return len(this.x)
}

// 出栈
func (this *stack) pop() {
	this.x = this.x[:len(this.x)-1]
}

// 返回栈顶元素
func (this *stack) top() int {
	return this.x[len(this.x)-1]
}