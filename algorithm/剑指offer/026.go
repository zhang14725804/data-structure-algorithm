/*
	补码：总和等于100000000000这种
	正数
	例1：1 补码 11**1111
	例2：2 补码 11**1110
	例3：3 补码 11**1101
	复数：绝对值的补码
	例1：-3 11**1101
	例2：-2 11**1110
*/
func NumberOf1(n int) int {
	// 处理负数
	if n < 0{
		n = pow(2,32) + n
	}
	fmt.Println(n)
	res := 0
	// todo：统计1的个数（位运算巧妙）
	for n > 0 {
		res += n & 1
		// 右移
		n = n >> 1
	}
	return res
}
func pow(x,y int) int {
	res := 1
	for y > 0 {
		res *= x
		y--
	}
	return res
}