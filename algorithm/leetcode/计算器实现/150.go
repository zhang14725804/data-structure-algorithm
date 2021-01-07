/*
	逆波兰表达式求值

	逆波兰表达式：

		逆波兰表达式是一种后缀表达式，所谓后缀就是指算符写在后面。

		平常使用的算式则是一种中缀表达式，如 ( 1 + 2 ) * ( 3 + 4 ) 。
		该算式的逆波兰表达式写法为 ( ( 1 2 + ) ( 3 4 + ) * ) 。
		逆波兰表达式主要有以下两个优点：

		去掉括号后表达式无歧义，上式即便写成 1 2 + 3 4 + * 也可以依据次序计算出正确结果。
		适合用栈操作运算：遇到数字则入栈；遇到算符则取出栈顶两个数字进行计算，并将结果压入栈中。
*/
func evalRPN(tokens []string) int {
	stk := &Stack{}

	for i := 0; i < len(tokens); i++ {
		// 判断是加减乘除还是数字；出栈计算式，注意顺序
		switch tokens[i] {
		case "+":
			op2 := stk.pop()
			op1 := stk.pop()
			stk.push(op1 + op2)
		case "-":
			op2 := stk.pop()
			op1 := stk.pop()
			stk.push(op1 - op2)
		case "*":
			op2 := stk.pop()
			op1 := stk.pop()
			stk.push(op1 * op2)
		case "/":
			op2 := stk.pop()
			op1 := stk.pop()
			stk.push(op1 / op2)
		default:
			// 入栈
			num, _ := strconv.Atoi(tokens[i])
			stk.push(num)
		}

	}
	return stk.pop()
}

type Stack struct {
	nums []int
}

func (this *Stack) push(n int) {
	this.nums = append(this.nums, n)
}
func (this *Stack) pop() int {
	res := this.nums[len(this.nums)-1]
	this.nums = this.nums[:len(this.nums)-1]
	return res
}