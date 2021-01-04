/*
	给定一个含有数字和运算符的字符串，为表达式添加括号，改变其运算优先级以求出不同的结果。你需要给出所有可能的组合的结果。有效的运算符号包含 +, - 以及 *

	思路( question )：
	（1） 不要思考整体，而是把目光聚焦局部，只看一个运算符。
	（2） 明确递归函数的定义是什么，相信并且利用好函数的定义。
*/

var memo map[string][]int = make(map[string][]int, 0) // 缓存消除重复计算
func diffWaysToCompute(input string) []int {
	if val, ok := memo[input]; ok {
		return val
	}
	res := make([]int, 0)
	// 扫描输入的input算式
	for i := 0; i < len(input); i++ {
		c := input[i]
		// 每当遇到运算符，进行分割
		if c == '-' || c == '*' || c == '+' {
			left := diffWaysToCompute(input[:i])
			right := diffWaysToCompute(input[i+1:])
			// 根据运算符合并结果
			for _, a := range left {
				for _, b := range right {
					if c == '+' {
						res = append(res, a+b)
					} else if c == '-' {
						res = append(res, a-b)
					} else if c == '*' {
						res = append(res, a*b)
					}
				}
			}
		}
	}
	// base case。如果 res 为空，说明算式是一个数字，没有运算符
	if len(res) == 0 {
		i, _ := strconv.Atoi(input)
		res = append(res, i)
	}
	memo[input] = res
	return res
}
