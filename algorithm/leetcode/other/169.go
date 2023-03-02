/*
	1. map记录元素出现的次数
	2. 遍历数组，统计次数
*/

func majorityElement(nums []int) int {
	nLen := len(nums)
	hash := make(map[int]int, 0)
	for _, v := range nums {
		hash[v]++
		if hash[v] > nLen/2 {
			return v
		}
	}
	return 0
}