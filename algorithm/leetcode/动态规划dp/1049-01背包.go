/*
	😅😅😅 本题其实就是尽量让石头分成重量相同的两堆，相撞之后剩下的石头最小，这样就化解成01背包问题了。
	本题物品的重量为store[i]，物品的价值也为store[i]。
	对应着01背包里的物品重量weight[i]和 物品价值value[i]。
*/

func lastStoneWeightII(stones []int) int {
    sum:=0
    for _,v:=range stones{
        sum+=v
    }
    target:=sum/2
	// （1）dp[j]表示容量（这里说容量更形象，其实就是重量）为j的背包，最多可以背dp[j]这么重的石头。
	// （3）因为重量都不会是负数，所以dp[j]都初始化为0就可以了
    dp:=make([]int,target+1)
	// 😅😅😅 （4）如果使用一维dp数组，物品遍历的for循环放在外层，遍历背包的for循环放在内层，且内层for循环倒叙遍历！
	// 遍历物品
    for i:=0;i<len(stones);i++{
		// 遍历背包
        for j:=target;j>=stones[i];j--{
			// （2）递推公式
            dp[j]=MaxInt(dp[j],dp[j-stones[i]]+stones[i])
        }
    }
	// dp[target]里是容量为target的背包所能背的最大重量
	// 那么分成两堆石头，一堆石头的总重量是dp[target]，另一堆就是sum - dp[target]
	// 在计算target的时候，target = sum / 2 因为是向下取整，所以sum - dp[target] 一定是大于等于dp[target]的
    return sum-dp[target]-dp[target]
}