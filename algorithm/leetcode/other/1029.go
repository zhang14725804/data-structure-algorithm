/*
	题意都没太看懂 😅
	1. 按照 price_A - price_B 从小到大排序；
	2. 将前 N 个人飞往 A 市，其余人飞往 B 市，并计算出总费用。
*/

func twoCitySchedCost(costs [][]int) int {
	// 😅
	sort.Slice(costs, func(i, j int) bool {
		return costs[i][0]-costs[i][1] < costs[j][0]-costs[j][1]
	})

	total := 0
	n := len(costs) / 2
	// 为了优化公司开支，发送前n个人到城市A，其他人去城市 B
	for i := 0; i < n; i++ {
		total += costs[i][0] + costs[i+n][1]
	}
	return total
}