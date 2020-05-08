/*
	数组中唯一只出现一次的数字
	二进制和异或的应用（难😅）
	todo：思路：状态机（更难，太难难了）
*/
func findNumberAppearingOnce(nums []int) int {
	ones,twos:=0,0
	for _,x := range nums{
		ones = (ones ^ x) & (~twos)
		twos = (twos ^ x) & (~ones)

	}
	return ones
}
