/*
	给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

	思路1：数学算法，秦九韶算法
	todo：1534236469测试未通过

	思路2：转换成字符串，reverse
*/

func reverse(x int) int {
	res := 0
	// 处理正负数
	pn := 1
	if x < 0 {
		pn = -1
		x = -x
	}
	// 重点在这里： x % 10 取余数，x /= 10取整
	for x > 0 {
		res = res*10 + x%10
		x /= 10
	}
	res *= pn
	if res < INT_MIN || res > INT_MAX {
		return 0
	}
	return res
}