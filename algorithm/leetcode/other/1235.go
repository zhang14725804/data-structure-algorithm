/*
	1. 将兼职时间按照endTime排序
	2. 用dp[i]表示当前i份工作可以获得的最大报酬
	3. 😅 状态转移方程 dp[i] = max(dp[i-1], dp[k]+profit[i-1])
	4. 【k】表示满足结束时间小于等于第 i−1 份工作开始时间的兼职工作数量
*/

func jobScheduling(startTime, endTime, profit []int) int {
	n := len(startTime)
	jobs := make([][3]int, n)
	// 😅 为什么要这么组织数据
	for i := 0; i < n; i++ {
		jobs[i] = [3]int{startTime[i], endTime[i], profit[i]}
	}
	sort.Slice(jobs, func(i, j int) bool { return jobs[i][1] < jobs[j][1] })

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		// 😅
		k := sort.Search(i, func(j int) bool { return jobs[j][1] > jobs[i-1][0] })
		dp[i] = max(dp[i-1], dp[k]+jobs[i-1][2])
	}
	return dp[n]
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
