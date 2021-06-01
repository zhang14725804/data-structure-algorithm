/*
	给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
*/

/*
	方法1：回溯算法

	怎么样写回溯算法(从上而下，※代表难点，根据题目而变化)：
	①画出递归树，找到状态变量(回溯函数的参数)，这一步非常重要※
	②根据题意，确立结束条件
	③找准选择列表(与函数参数相关),与第一步紧密关联※
	④判断是否需要剪枝
	⑤作出选择，递归调用，进入下一层
	⑥撤销选择

	没想出来终止条件 😅😅😅
*/
var ans [][]int
var path []int

func subsets(nums []int) [][]int {
	backtrack(nums, 0)
	ans = make([][]int, 0) // 只是为了提交，leetcode提交时，ans 会拼接之前提交的结果
	return ans
}

func backtrack(nums []int, start int) {

	// 收集子集，要放在终止添加的上面，否则会漏掉自己
	back := make([]int, len(path))
	copy(back, path)
	ans = append(ans, back)

	// 😅 无需判断结束条件，终止条件可以不加
	if start >= len(nums) {
		return
	}
	for i := start; i < len(nums); i++ {
		// 选择
		path = append(path, nums[i])
		// 回溯； i+1 而不是start+1
		backtrack(nums, i+1)
		// 撤销选择
		path = path[:len(path)-1]
	}
}