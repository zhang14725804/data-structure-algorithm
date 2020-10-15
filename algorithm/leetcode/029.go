/*
	给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

	返回被除数 dividend 除以除数 divisor 得到的商。

	整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2

	-2147483648,-1 测试不通过（golang中int，在32 位系统下是32 位，而在64 位系统下是64 位）
	其数值范围是 [−231,  231 − 1]

*/
const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

func divide(dividend int, divisor int) int {
	// 判断0，+-1的情况
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}

	if divisor == -1 {
		//
		if dividend > INT_MIN {
			return -dividend
		}
		return INT_MAX
	}

	// 判断符号
	sign := 1
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		sign = -1
	}
	// 修正符号
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	res := div(dividend, divisor)
	// 处理返回结果
	if sign > 0 {
		if res > INT_MAX {
			return INT_MAX
		}
		return res
	}
	return -res
}

func div(a, b int) int {
	if a < b {
		return 0
	}
	count := 1
	tb := b
	// 做除法（question）
	for tb+tb <= a {
		count += count
		tb += tb
	}
	return count + div(a-tb, b)
}