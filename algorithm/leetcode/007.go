/*
	给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

	思路1：数学算法，秦九韶算法
	todo：1534236469测试未通过

	思路2：转换成字符串，reverse
*/

const INT_MAX = (1 << 31) - 1
const INT_MIN = -(1 << 31)

func reverse(x int) int {
	res := 0
	pn := 1 // 正负数
	if x < 0 {
		pn = -1
		x = -x
	}
	// 每次遍历一位
	for x > 0 {
		res = res*10 + x%10
		x /= 10
	}
	res *= pn
	// 处理极值
	if res < INT_MIN || res > INT_MAX {
		return 0
	}
	return res
}