/*
	46. Permutations

	全排列问题，两种方法（这特么。。。）
	（1）枚举每个位置上放那个树
	（2）枚举每个数放在哪个位置

	todos：：dfs递归出问题，导致程序有问题(没看出问题出在哪里)
*/

var (
	n    int // 数组大小
	st   []bool
	ans  [][]int // 所有方法
	path []int   // 当前方案
)

func permute(nums []int) [][]int {
	// 方法一：枚举每个位置放什么数字
	n = len(nums)
	//
	st = make([]bool, n)
	dfs(nums, 0)
	return ans
}

func dfs(nums []int, u int) {
	// 已经便利完所有方案
	if u == n {
		ans = append(ans, path)
		return
	}
	for i := 0; i < n; i++ {
		// 当前位置可以放这个数
		if !st[i] {
			st[i] = true
			path = append(path, nums[i])
			// 递归下一层
			dfs(nums, u+1)
			path = path[:len(path)-1]
			st[i] = false
		}
	}
}