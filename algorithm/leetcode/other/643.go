
/*
	1. 先算出前4个数之和
	2. 循环数组，更新最大和（这里只需要拿掉第一个加上最后一个，而不是每次都算一遍所有的和）
	3. 只需要每次比较sum，计算一次res即可
*/
func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for _, v := range nums[:k] {
		sum += v
	}
	maxSum := sum
	for i := k; i < len(nums); i++ {
		sum = sum - nums[i-k] + nums[i]
		maxSum = max(maxSum, sum)
	}
	return float64(maxSum) / float64(k)
}

/*
	我的实现不足之处 😅😅😅：
		1. 每次都sum一次
		2. 每次都float64类型转换之后，更新一次res
*/
func findMaxAverage(nums []int, k int) float64 {
	var res float64
	for i := 0; i+k < len(nums); i++ {
		cres := float64(sum(nums[:i+k])) / float64(k)
		if cres > res {
			res = cres
		}
	}
	return res
}