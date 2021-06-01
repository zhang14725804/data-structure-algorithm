/*
	给定一个整型数组, 你的任务是找到所有该数组的递增子序列，递增子序列的长度至少是 2 。
	给定数组的长度不会超过15。
	数组中的整数范围是 [-100,100]。
	给定数组中可能包含重复数字，相等的数字应该被视为递增的一种情况。
*/

/*
	方法1：回溯
	 😅 而本题求自增子序列，是不能对原数组经行排序的，排完序的数组都是自增子序列了

	😅😅😅 可优化的点：当前使用map来记录本层元素是否重复使用，数值范围[-100,100]，所以完全可以用【数组】来做哈希
*/
var ans [][]int
var path []int
var nums []int

func findSubsequences(_nums []int) [][]int {
	nums = _nums
	ans = make([][]int, 0) // 只是为了提交，leetcode提交时，ans 会拼接之前提交的结果
	backtrack(0)
	return ans
}
func backtrack(start int) {
	// 😅 注意这里不要加return，要取树上的节点
	if len(path) >= 2 {
		back := make([]int, len(path))
		copy(back, path)
		ans = append(ans, back)
	}
	// 😅😅😅 可优化的地方 used := make([]bool, 201) // 这里使用数组来进行去重操作，题目说数值范围[-100, 100]
	// 记录本层元素是否重复使用，新的一层useMap都会重新定义（清空），所以要知道usedMap只负责【本层】！
	usedMap := make(map[int]bool, 0)
	for i := start; i < len(nums); i++ {
		// 😅 本层未使用过，且 符合递增条件
		if len(path) > 0 && nums[i] < path[len(path)-1] || usedMap[nums[i]] {
			continue
		}
		// 记录这个元素在本层用过了，本层后面不能再用了
		usedMap[nums[i]] = true
		path = append(path, nums[i])
		backtrack(i + 1)
		path = path[:len(path)-1]
		// 😅 这里不需要 usedMap[nums[i]] = false
	}
}