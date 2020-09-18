/*
	编写一个函数，输入是一个无符号整数，返回其二进制表达式中数字位数为 ‘1’ 的个数（也被称为汉明重量）。
	todo
*/
func hammingWeight(num uint32) int {
	nums := strconv.FormatUint(uint64(num), 2)
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == '1' {
			res++
		}
	}
	return res
}