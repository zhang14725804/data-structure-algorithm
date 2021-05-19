## 回溯算法

DFS 是一个劲的往某一个方向搜索，而回溯算法建立在 DFS 基础之上的，但不同的是在搜索过程中，达到结束条件后，恢复状态，回溯上一层，再次搜索。因此回溯算法与 DFS 的区别就是有无状态重置

只需要思考 3 个问题：

	1、路径：也就是已经做出的选择。
	2、选择列表：也就是你当前可以做的选择。
	3、结束条件：也就是到达决策树底层，无法再做选择的条件。

其核心就是 for 循环里面的递归，在递归调用之前【做选择】，在递归调用之后【撤销选择】

```golang
// 路径，选择列表
func backtrack(nums []int, u int) {
	// 结束条件
	if u == n {
		c := make([]int, n)
		copy(c, path)
		ans = append(ans, c)
		return
    }
    
	for i := 0; i < n; i++ {
        // 做选择
        st[i] = true
        path = append(path, nums[i])
        // 递归下一层
        backtrack(nums, u+1)
        // 撤销选择
        st[i] = false
        path = path[:len(path)-1]
	}
}
```

- 具体题目

        22.括号生成
        37.解数独
        77.组合
        78.子集
        46.全排列问题
        51.n皇后问题

todo:还可以通过剪枝进行优化