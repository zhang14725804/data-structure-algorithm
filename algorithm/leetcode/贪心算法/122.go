/*
	（1）只有一只股票！
	（2）当前只有买股票或者买股票的操作

	其实我们需要收集每天的正利润就可以，收集正利润的区间，就是股票买卖的区间，而我们只需要关注最终利润，不需要记录区间。
	那么只收集正利润就是贪心所贪的地方！
	局部最优：收集每天的正利润，全局最优：求得最大利润。
*/

func maxProfit(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		res += MaxInt(prices[i]-prices[i-1], 0)
	}
	return res
}