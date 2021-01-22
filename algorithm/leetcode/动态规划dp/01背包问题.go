/*
	给你一个可装载重量为W的背包和N个物品，每个物品有重量和价值两个属性。
	其中第i个物品的重量为wt[i]，价值为val[i]，现在让你用这个背包装物品，最多能装的价值是多少？

	N = 3, W = 4
	wt = [2, 1, 3]
	val = [4, 2, 3]
	返回 6，选择前两件物品装进背包，总重量 3 小于W，可以获得最大价值 6
*/

/*
	动态规划思路( question )：
	第一步要明确两点，「状态」和「选择」
	（1）状态有两个，就是「背包的容量」和「可选择的物品」
	（2）选择就是「装进背包」或者「不装进背包」
	第二步要明确dp数组的定义
	dp[i][w]的定义如下：对于前i个物品，当前背包的容量为w，这种情况下可以装的最大价值是dp[i][w]
	第三步，根据「选择」，思考状态转移的逻辑
	（1）不选择第i个物品装入背包: dp[i][w]=dp[i-1][w]
	（2）选第i个物品装入了背包: dp[i][w]=dp[i-1][w-wt[i-1]] + val[i-1]
*/
func KnapSack01(weights []int, values []int, N int, W int) int {
	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, W+1)
	}
	// 枚举每一种状态
	for i := 1; i <= N; i++ {
		for j := 1; j <= W; j++ {
			// 当前背包容量装不下，只能选择不装入
			if j-weights[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				// 装入或者不装入
				dp[i][j] = MaxInt(dp[i-1][j-weights[i-1]]+values[i-1], dp[i-1][j])
			}
		}
	}
	return dp[N][W]
}
