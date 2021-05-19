/*
	47. Permutations II
	回溯算法（todo）
	给定一个可包含重复数字的序列，返回所有不重复的全排列

	全排列问题，两种方法（这特么。。。）
	（1）枚举每个位置上放那个树
	（2）枚举每个数放在哪个位置
*/
/*
	方法二：枚举每个数放在哪个位置（这个题需要判重）

	将所有相同的数放在一起 sort
	认为规定相同数字的相对顺序：不变
	dfs增加状态
*/
var (
	n    int // 数组大小
	st   []bool
	ans  [][]int // 所有方法
	path []int   // 当前方案
)

func permuteUnique(nums []int) [][]int {
	n = len(nums)
	st = make([]bool, n)
	path = make([]int, n)
	// 排序
	nums = quickSort(nums)
	dfs(nums, 0, 0)
	return ans
}

func dfs(nums []int, u, start int) {
	if u == n {
		// 这里是个什么坑（question）。直接ans = append(ans, path), [[1,1,2],[1,1,2],[1,1,2]]
		c := make([]int, len(path))
		copy(c, path)
		ans = append(ans, c)
		return
	}
	for i := start; i < n; i++ {
		if !st[i] {
			st[i] = true
			path[i] = nums[u]
			var s int
			if u+1 < n && nums[u+1] == nums[u] {
				s = i + 1
			} else {
				s = 0
			}
			dfs(nums, u+1, s)
			st[i] = false
		}
	}
}