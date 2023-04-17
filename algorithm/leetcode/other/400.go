/*
	TODO
*/
func findNthDigit(n int) int {
	// 当前遍历到的位数
	d := 1
	// count 当前位数下的所有整数的位数之和
	for count := 9; n > d*count; count *= 10 {
		n -= d * count
		d++
	}
	// 下标从0开始
	idx := n - 1
	// 得到无限整数序列中的第 n 位数字 😅😅😅 不懂
	start := int(math.Pow10(d - 1))
	num := start + idx/d
	didx := idx % d
	res := num / int(math.Pow10(d-didx-1)) % 10
	return res
}