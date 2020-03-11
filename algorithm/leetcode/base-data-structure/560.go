/*
	560. Subarray Sum Equals K

	考察的算法：前缀和和hash（边界问题）
*/
func subarraySum(nums []int, k int) int {
	hash := make(map[int]int,0)
	// 注意边界问题
	hash[0] = 1
	res := 0
	for i,sum:=0,0;i<len(nums);i++{
		// 前缀和
		sum += nums[i]
		res+=hash[sum-k]
		hash[sum]+=1
	}
	return res
}