/*
	给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。

	思路1：hash
	思路2：位操作 一通骚操作  😅😅😅
	思路3：todo:困难（https://leetcode.com/problems/single-number-ii/discuss/43295/Detailed-explanation-and-generalization-of-the-bitwise-operation-method-for-single-numbers）
*/
func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < 64; i++ {
		sum := 0
		// 统计1的个数 😅😅😅
		for j := 0; j < len(nums); j++ {
			sum += (nums[j] >> i) & 1
		}
		// 还原  😅😅😅
		result ^= (sum % 3) << i
	}
	return result
}
