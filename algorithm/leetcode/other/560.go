/*
	前缀和 为啥我想不到 😅
*/
func subarraySum1(nums []int, k int) int {
	size := len(nums)
	// 😅 前缀和数组，为什么要【size+1】
	preSum := make([]int, size+1)
	preSum[0] = 0
	// 构造前缀和
	for i := 0; i < size; i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}
	ans := 0
	for i := 1; i <= size; i++ {
		for j := 0; j < i; j++ {
			// 😅😅😅 nums[j..i-1]区间和
			if preSum[i]-preSum[j] == k {
				ans++
			}
		}
	}
	return ans
}

/*
	前缀和和hash（边界问题）
	😅😅😅 不懂
*/
func subarraySum(nums []int, k int) int {
	// key：前缀和，value：key 对应的前缀和的个数
	hash := make(map[int]int, 0)
	// 😅 对于下标为 0 的元素，前缀和为 0，个数为 1
	hash[0] = 1
	res := 0
	//  [0..i][0..i] 里所有数的和
	preSum := 0
	for i := 0; i < len(nums); i++ {
		// 前缀和
		preSum += nums[i]
		// 先获得前缀和为 preSum - k 的个数，加到计数变量里
		res += hash[preSum-k]
		// 然后维护 preSum 的定义
		hash[preSum] += 1
	}
	return res
}