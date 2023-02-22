/*
	第k个排列（全排列问题 question）


*/

/*
	方法1：深度优先遍历 + 剪枝、有序数组模拟（question，没太懂😅）

	所求排列 一定在叶子结点处得到，进入每一个分支，可以根据已经选定的数的个数，进而计算还未选定的数的个数，然后计算阶乘，就知道这一个分支的 叶子结点 的个数：
	如果 kk 大于这一个分支将要产生的叶子结点数，直接跳过这个分支，这个操作叫「剪枝」；
	如果 kk 小于等于这一个分支将要产生的叶子结点数，那说明所求的全排列一定在这一个分支将要产生的叶子结点里，需要递归求解。
*/
var (
	used      []bool // 记录数字是否使用过
	factorial []int  // 阶乘数数组
	n, k      int
	path      string
)

func getPermutation(_n int, _k int) string {
	n, k = _n, _k
	calculateFactorial(n)
	used = make([]bool, n+1)
	dfs(0)

	return path
}

//  index 在这一步之前已经选择了几个数字，其值恰好等于这一步需要确定的下标位置
func dfs(index int) {
	if index == n {
		return
	}
	// 计算还未确定的数字的全排列的个数，第 1 次进入的时候是 n - 1
	cnt := factorial[n-1-index]
	for i := 1; i <= n; i++ {
		if used[i] {
			continue
		}
		if cnt < k {
			k -= cnt
			continue
		}
		path += fmt.Sprint(i)
		used[i] = true
		dfs(index + 1)
		return
	}
}
func calculateFactorial(n int) {
	factorial = make([]int, n+1)
	factorial[0] = 1
	for i := 1; i <= n; i++ {
		factorial[i] = factorial[i-1] * i
	}
}

/*
	方法二：有序数组（链表）模拟（question）
*/