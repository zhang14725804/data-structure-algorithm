/*
	给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。你可以按 任意顺序 返回答案。
	进阶：你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现？

	思路1：hash+for
	思路2：位运算  genius 😅😅😅😅😅😅,思路妙啊
*/
func singleNumber(nums []int) []int {
	// 所有的数字异或（^） 得到 first^second的结果
	diff := 0
	for _, num := range nums {
		diff ^= num
	}
	// 得到只有一位 1 的数
	diff2 := diff & (-diff)
	first, second := 0, 0
	// 将原来的数分为两组，一组在某个位置 都是1，另一组在此位置都为0
	for _, num := range nums {
		if diff2&num == 0 {
			first ^= num
		}
	}
	// first^second^first = second
	second = diff ^ first
	return []int{first, second}
}