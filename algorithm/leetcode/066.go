/*
	除 99 之外的数字加一；
	数字 99。
*/
func plusOne(digits []int) []int {
	n := len(digits)
	// 妙啊
	for i := n - 1; i >= 0; i-- {
		digits[i]++
		digits[i] = digits[i] % 10 // 判断最后一位
		if digits[i] != 0 {
			return digits
		}
	}
	// 99，999的情况
	digits = make([]int, n+1)
	digits[0] = 1
	return digits
}