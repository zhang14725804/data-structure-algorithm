/*
	模拟手算乘法的过程
	给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

	核心：num1[i]和num2[j]的乘积对应的就是res[i+j]和res[i+j+1]这两个位置
*/
func multiply(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	// 结果最多m+n位
	res := make([]int, n+m)
	// 从个位数开始逐位相乘
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			// 相乘
			pro := int(num1[i]-'0') * int(num2[j]-'0')
			// 乘积在res对应的索引位置
			p1 := i + j
			p2 := i + j + 1
			// 叠加到res上
			sum := res[p2] + pro
			res[p2] = sum % 10
			res[p1] += sum / 10
		}
	}
	// 排除未使用的位，前缀可能位0
	i := 0
	for i < len(res) && res[i] == 0 {
		i++
	}
	// 转为字符串
	str := ""
	for ; i < len(res); i++ {
		// int转string
		str += fmt.Sprint(res[i])
	}
	if len(str) == 0 {
		return "0"
	}
	return str
}