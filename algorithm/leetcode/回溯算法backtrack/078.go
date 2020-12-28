/*
	给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

	迭代，回溯，位操作三种方法。位操作方法最骚(todo)

	怎么样写回溯算法(从上而下，※代表难点，根据题目而变化)：
	①画出递归树，找到状态变量(回溯函数的参数)，这一步非常重要※
	②根据题意，确立结束条件
	③找准选择列表(与函数参数相关),与第一步紧密关联※
	④判断是否需要剪枝
	⑤作出选择，递归调用，进入下一层
	⑥撤销选择
*/
var ans [][]int
var path []int

func subsets(nums []int) [][]int {
	backtrack(nums, 0)
	return ans
}

// question 地址引用有问题
func backtrack(nums []int, start int) {
	temp := make([]int, len(path))
	copy(temp, path)
	ans = append(ans, temp)
	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		backtrack(nums, i+1)
		path = path[:len(path)-1]
	}
}