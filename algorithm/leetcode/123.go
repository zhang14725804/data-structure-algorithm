/*
	你最多可以完成 两笔 交易。你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

	思路：前后缀分解
*/
func maxProfit(prices []int) int {
	n:=len(prices)

	f:=make([]int,n+2) // 这里slice长度为什么n+2
	minp:=prices[0]
	// 前缀最大值
	for i:=1;i<=n;i++{
		f[i] = compare(f[i-1],prices[i-1]-minp,true)
		minp = compare(minp,prices[i-1],false)
	}
	// todo：代码不理解
	res := 0
	maxp := 0
	// 后缀和
	for i:=n;i>0;i--{
		res = compare(res,maxp-prices[i-1]+f[i-1],true)
		maxp = compare(maxp,prices[i-1],true)
	}

	return res

}