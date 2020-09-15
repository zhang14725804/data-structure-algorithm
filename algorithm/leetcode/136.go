/*
	给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

	思路1：hash
	思路2：所有出现的数字乘以2，再减去所有的数字之和
	思路3：异或
*/
func singleNumber(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum ^= num
	}
	return sum
}
