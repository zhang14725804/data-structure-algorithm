/*
	560. Subarray Sum Equals K

	把这个题改成多少和小于等于K：：平衡树的基础上二分
*/

/*
	前缀和
*/
func subarraySum1(nums []int, k int) int {
	size := len(nums)
	// 前缀和数组
	preSum := make([]int, size+1)
	preSum[0] = 0
	// 构造前缀和
	for i := 0; i < size; i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}
	ans := 0
	for i := 1; i <= size; i++ {
		for j := 0; j < i; j++ {
			// sum of nums[j..i-1]
			if preSum[i]-preSum[j] == k {
				ans++
			}
		}
	}
	return ans
}

/*
	前缀和和hash（边界问题）
*/
func subarraySum(nums []int, k int) int {
	// 以和为键，出现次数为对应的值
	hash := make(map[int]int, 0)
	// 注意边界问题
	hash[0] = 1
	res := 0
	//  [0..i][0..i] 里所有数的和
	preSum := 0
	for i := 0; i < len(nums); i++ {
		// 前缀和
		preSum += nums[i]
		// question 不理解
		res += hash[preSum-k]
		hash[preSum] += 1
	}
	return res
}