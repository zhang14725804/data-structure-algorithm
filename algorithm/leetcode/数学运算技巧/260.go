/*
	思路1：hash+for （我想到的这一种😅）
	思路2：位运算  genius 😅😅😅😅😅😅
*/
func singleNumber(nums []int) []int {
	// x1和x2 异或的和，其他的数字因为出现两次被抵消了
	xor := 0
	for _, num := range nums {
		xor ^= num
	}

	// 😅😅😅 取最低有效位
	lsb := xor & -xor
	// 用最低有效位把所有数字分为两类
	tp1, tp2 := 0, 0
	for _, num := range nums {
		if num&lsb > 0 {
			tp1 ^= num
		} else {
			tp2 ^= num
		}
	}
	return []int{tp1, tp2}
}