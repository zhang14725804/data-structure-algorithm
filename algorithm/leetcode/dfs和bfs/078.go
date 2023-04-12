/*
	方法1: dfs
*/
func subsets(nums []int) (ans [][]int) {
	set := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			// 😅😅😅
			ans = append(ans, append([]int{}, set...))
			return
		}

		// 😅 考虑选择当前位置
		set = append(set, nums[cur])
		dfs(cur + 1)
		// 😅 考虑不选择当前位置
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return
}

/*
	方法2: 二进制思想（这个思路不错 😅😅😅）
	1. 数组的每个元素，可以有两个状态，在子数组中和不在子数组中，所有状态的组合就是所有子数组了
	2. 数组长度是 n，那么每个比特位是 2 个状态，所有总共就是 2 的 n 次方个子数组
*/
func subsets(nums []int) [][]int {
	var res [][]int
	//
	for i := 0; i < 1<<len(nums); i++ {
		var now []int
		// 判断每个元素
		for j := 0; j < len(nums); j++ {
			// 判断每个比特位是否是 1，是 1 的话将对应数字加入即可
			if i>>j&1 == 1 {
				now = append(now, nums[j])
			}
		}
		res = append(res, now)
	}
	return res
}