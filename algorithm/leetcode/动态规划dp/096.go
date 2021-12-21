/*
	动态规划
*/
func numTrees(n int) int {
	// dp[i] ： 1到i为节点组成的二叉搜索树的个数为dp[i]。
	dp := make([]int, n+1)
	// dp数组初始化
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			// 😅 递推公式
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

/*
	递归 😅😅😅
	我们可以利用一下查找二叉树的性质。左子树的所有值小于根节点，右子树的所有值大于根节点。

	所以如果求 1...n 的所有可能。

	我们只需要把 1 作为根节点，[ ] 空作为左子树，[ 2 ... n ] 的所有可能作为右子树。

	2 作为根节点，[ 1 ] 作为左子树，[ 3...n ] 的所有可能作为右子树。

	3 作为根节点，[ 1 2 ] 的所有可能作为左子树，[ 4 ... n ] 的所有可能作为右子树，然后左子树和右子树两两组合。

	4 作为根节点，[ 1 2 3 ] 的所有可能作为左子树，[ 5 ... n ] 的所有可能作为右子树，然后左子树和右子树两两组合。

	...

	n 作为根节点，[ 1... n ] 的所有可能作为左子树，[ ] 作为右子树。

	至于，[ 2 ... n ] 的所有可能以及 [ 4 ... n ] 以及其他情况的所有可能，可以利用上边的方法，把每个数字作为根节点，然后把所有可能的左子树和右子树组合起来即可。

	如果只有一个数字，那么所有可能就是一种情况，把该数字作为一棵树。而如果是 [ ]，那就返回 null。
*/
var memoization = make(map[int]int, 0)

func numTrees(n int) int {
	if n == 0 {
		return 0
	}
	return dfs(n)
}

func dfs(n int) int {
	if val, ok := memoization[n]; ok {
		return val
	}
	// base case 此时没有数字或者只有一个数字,返回 1
	if n == 0 || n == 1 {
		return 1
	}
	ans := 0
	for i := 1; i <= n; i++ {
		// 得到所有可能的左子树
		left := dfs(i - 1)
		// 得到所有可能的右子树
		right := dfs(n - i)
		// 左子树右子树两两组合
		ans += left * right
	}
	memoization[n] = ans
	return ans
}