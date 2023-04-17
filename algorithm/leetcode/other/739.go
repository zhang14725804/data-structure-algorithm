/*
	我的笨办法 😅
*/
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, 0)
	// 遍历每一天
	for i := 0; i < len(temperatures); i++ {
		// 默认0
		idx := 0
		// 遍历后续的每一天
		for j := i + 1; j < len(temperatures); j++ {
			// 符合条件的日子
			if temperatures[j] > temperatures[i] {
				idx = j - i
				break
			}
		}
		res = append(res, idx)
	}
	return res
}

/*
	单调栈 😅😅😅
	维护一个存储下标的单调栈，从栈底到栈顶的下标对应的温度列表中的温度依次递减。
	如果一个下标在单调栈里，则表示尚未找到下一次温度更高的下标。
*/
func dailyTemperatures(temperatures []int) []int {
	tLen := len(temperatures)
	res := make([]int, tLen)
	stack := []int{}
	for i := 0; i < tLen; i++ {
		tp := temperatures[i]
		// 当前温度大于栈顶温度
		for len(stack) > 0 && tp > temperatures[stack[len(stack)-1]] {
			// 出栈
			prevIdx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 更新res值
			res[prevIdx] = i - prevIdx
		}
		stack = append(stack, i)
	}
	return res
}