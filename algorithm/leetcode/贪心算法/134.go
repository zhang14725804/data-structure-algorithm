/*
	在一条环路上有 N 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
	你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。
	你从其中的一个加油站出发，开始时油箱为空。
	如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。
*/

/*
	思路(思路精巧)：从i个加油站开始循环，那个加油站为-1了，排除之前的加油站，再从当前加油站重新开始寻找

	局部最优：当前累加rest[j]的和curSum一旦小于0，起始位置至少要是j+1，因为从j开始一定不行。全局最优：找到可以跑一圈的起始位置。
*/
func canCompleteCircuit(gas []int, cost []int) int {
	// totalGas totalCost
	tg, tc := 0, 0
	for i := 0; i < len(gas); i++ {
		tc += cost[i]
		tg += gas[i]
	}
	if tg < tc {
		return -1
	}
	// currentGas,start
	// 😅😅 两个for循环有些蠢
	cg, s := 0, 0
	for i := 0; i < len(gas); i++ {
		cg += gas[i] - cost[i]
		if cg < 0 {
			cg = 0
			s = i + 1
		}
	}
	return s
}

func canCompleteCircuit(gas []int, cost []int) int {
	curSum := 0
	totalSum := 0
	start := 0
	// 😅 一个for循环搞定
	for i := 0; i < len(gas); i++ {
		curSum += gas[i] - cost[i]
		totalSum += gas[i] - cost[i]
		// 当前累加rest[i]和 curSum一旦小于0
		if curSum < 0 {
			// 起始位置更新为i+1
			start = i + 1
			// curSum从0开始
			curSum = 0
		}
	}
	// 说明怎么走都不可能跑一圈了
	if totalSum < 0 {
		return -1
	}
	return start
}