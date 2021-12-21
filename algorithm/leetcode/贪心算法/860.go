/*
	有如下三种情况：

	情况一：账单是5，直接收下。
	情况二：账单是10，消耗一个5，增加一个10
	情况三：账单是20，优先消耗一个10和一个5，如果不够，再消耗三个5

	此时大家就发现 情况一，情况二，都是固定策略，都不用我们来做分析了，而唯一不确定的其实在情况三。

	而情况三逻辑也不复杂甚至感觉纯模拟就可以了，其实情况三这里是有贪心的。

	账单是20的情况，为什么要优先消耗一个10和一个5呢？

	因为美元10只能给账单20找零，而美元5可以给账单10和账单20找零，美元5更万能！

	所以局部最优：遇到账单20，优先消耗美元10，完成本次找零。全局最优：完成全部账单的找零。

	局部最优可以推出全局最优，并找不出反例，那么就试试贪心算法！
*/
func lemonadeChange(bills []int) bool {
	m5, m10, m20 := 0, 0, 0
	for _, val := range bills {
		// case1
		if val == 5 {
			m5++
		}
		// case2
		if val == 10 {
			if m5 < 1 {
				return false
			}
			m10++
			m5--
		}
		// case3
		if val == 20 {
			if m5 > 0 && m10 > 0 {
				m10--
				m5--
				m20++
			} else if m5 >= 3 {
				m5 -= 3
				m20++
			} else {
				return false
			}
		}
	}
	return true
}