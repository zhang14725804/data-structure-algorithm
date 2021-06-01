/*
	输入一组数字（可能包含重复数字），输出其所有的排列方式
	todo：暴力搜索问题
	（1）排序（目的是把相同的数字找出来）
	（2）选择位置
	todo：代码有点难理解,思路倒是有(或者也可以暴力枚举，然后去重)
*/

var ans [][]int
var path []int

func permutation(nums []int) [][]int {
	//  non-constant array bound len(nums)
	// path = [len(nums)][]int😅
	// 初始化数组
	path = make([]int, len(nums))
	// 排序
	BubbleSort(nums)
	dfs(nums, 0, 0, 0)
	return ans
}

// todo：用二进制骚操作
func dfs(nums []int, u, start, state int) {
	if u == len(nums) {
		ans = append(ans, path)
		return
	}
	if u == 0 || nums[u] != nums[u-1] {
		start = 0
	}
	for i := start; i < len(nums); i++ {
		if (state >> i & 1) == 0 {
			path[i] = nums[u]
			dfs(nums, u+1, i+1, state+(1<<i))
		}
	}
}