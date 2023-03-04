/*
	方法1：hash
	1. map统计字符出现的次数
	2. 循环map，找到符合条件的数字
*/
func majorityElement(nums []int) []int {
	cnt := map[int]int{}
	for _, n := range nums {
		cnt[n]++
	}
	res := make([]int, 0)
	for n, c := range cnt {
		if c > len(nums)/3 {
			res = append(res, n)
		}
	}
	return res
}

/*
	方法2:摩尔投票法
*/