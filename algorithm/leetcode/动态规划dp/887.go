/*
	你将获得 K 个鸡蛋，并可以使用一栋从 1 到 N  共有 N 层楼的建筑。

	每个蛋的功能都是一样的，如果一个蛋碎了，你就不能再把它掉下去。

	你知道存在楼层 F ，满足 0 <= F <= N 任何从高于 F 的楼层落下的鸡蛋都会碎，从 F 楼层或比它低的楼层落下的鸡蛋都不会破。

	每次移动，你可以取一个鸡蛋（如果你有完整的鸡蛋）并把它从任一楼层 X 扔下（满足 1 <= X <= N）。

	你的目标是确切地知道 F 的值是多少。

	无论 F 的初始值如何，你确定 F 的值的最小移动次数是多少？
*/

/*
	动态规划思路( question )：
	第一步要明确两点，「状态」和「选择」
	「状态」：当前拥有的鸡蛋数K和需要测试的楼层数N
	「选择」：去选择哪层楼扔鸡蛋

	第二步要明确【dp数组】者带有【两个状态参数的dp函数】的定义

	第三步，根据「选择」，思考状态转移的逻辑
	如果鸡蛋碎了，那么鸡蛋的个数K应该减一，搜索的楼层区间应该从[1..N]变为[1..i-1]共i-1层楼；
	如果鸡蛋没碎，那么鸡蛋的个数K不变，搜索的楼层区间应该从 [1..N]变为[i+1..N]共N-i层楼

	TODO:代码有问题
*/
func superEggDrop(K int, N int) int {
	memo := make(map[string]int)
	var dp func(K, N int) int

	dp = func(K, N int) int {
		// base case
		// 当楼层数N等于 0 时，显然不需要扔鸡蛋；当鸡蛋数K为 1 时，显然只能线性扫描所有楼层
		if K == 1 {
			return N
		}
		if N == 0 {
			return 0
		}
		// 避免重复计算
		key := fmt.Sprintf("%v", N) + fmt.Sprintf("%v", K)
		if value, ok := memo[key]; ok {
			return value
		}
		// 举所有可能的选择
		res := (1 << 32)
		for i := 1; i < N+1; i++ {
			// dp(K - 1, i - 1), # 碎
			// dp(K, N - i)      # 没碎
			// + 1 # 在第 i 楼扔了一次
			// 最坏情况下的最少扔鸡蛋次数
			res = MinInt(res, MaxInt(dp(K, N-i), dp(K-1, i-1))+1)
		}
		memo[key] = res
		return res
	}
	return dp(N, K)
}