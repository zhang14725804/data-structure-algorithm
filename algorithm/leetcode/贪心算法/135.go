/*
	老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。

	你需要按照以下要求，帮助老师给这些孩子分发糖果：

	每个孩子至少分配到 1 个糖果。
	相邻的孩子中，评分高的孩子必须获得更多的糖果。

	那么这样下来，老师至少需要准备多少颗糖果呢？

	如果当前小朋友的 rating 比后一个小朋友的大，那么理论上当前小朋友的糖要比后一个的小朋友的糖多，但此时后一个小朋友的糖还没有确定，怎么办呢？

*/

/*
	这道题目一定是要确定一边之后，再确定另一边，例如比较每一个孩子的左边，然后再比较右边，【1如果两边一起考虑一定会顾此失彼】。

	先确定右边评分大于左边的情况（也就是【从前向后遍历】）

	此时局部最优：只要右边评分比左边大，右边的孩子就多一个糖果，全局最优：相邻的孩子中，评分高的右孩子获得比左边孩子更多的糖果

	局部最优可以推出全局最优。

	如果ratings[i] > ratings[i - 1] 那么[i]的糖 一定要比[i - 1]的糖多一个，所以贪心：candyVec[i] = candyVec[i - 1] + 1

	再确定左孩子大于右孩子的情况（【从后向前遍历】）

	遍历顺序这里有同学可能会有疑问，为什么不能从前向后遍历呢？

	因为如果从前向后遍历，根据 ratings[i + 1] 来确定 ratings[i] 对应的糖果，那么每次都不能利用上前一次的比较结果了。

	所以确定左孩子大于右孩子的情况一定要从后向前遍历！

	如果 ratings[i] > ratings[i + 1]，此时candyVec[i]（第i个小孩的糖果数量）就有两个选择了，一个是candyVec[i + 1] + 1（从右边这个加1得到的糖果数量），一个是candyVec[i]（之前比较右孩子大于左孩子得到的糖果数量）。

	那么又要贪心了，局部最优：取candyVec[i + 1] + 1 和 candyVec[i] 最大的糖果数量，保证第i个小孩的糖果数量即大于左边的也大于右边的。全局最优：相邻的孩子中，评分高的孩子获得更多的糖果。

	局部最优可以推出全局最优。

	所以就取candyVec[i + 1] + 1 和 candyVec[i] 最大的糖果数量，candyVec[i]只有取最大的才能既保持对左边candyVec[i - 1]的糖果多，也比右边candyVec[i + 1]的糖果多。


	采用了两次贪心的策略：
	一次是从左到右遍历，只比较右边孩子评分比左边大的情况。
	一次是从右到左遍历，只比较左边孩子评分比右边大的情况。
*/
func candy(ratings []int) int {
	n := len(ratings)
	candies := make([]int, n)

	// 每人先发一个糖
	for i := 0; i < n; i++ {
		candies[i] = 1
	}

	// 😅 正着进行
	for i := 0; i < n-1; i++ {
		if ratings[i] < ratings[i+1] {
			candies[i+1] = candies[i] + 1
		}
	}

	// 😅 倒着进行
	for i := n - 1; i > 0; i-- {
		if ratings[i] < ratings[i-1] {
			// 后一个小朋友的糖果没有前一个的多，就更新后一个等于前一个加 1
			if candies[i-1] <= candies[i] {
				candies[i-1] = candies[i] + 1
			}
		}
	}

	// 求和
	sum := 0
	for i := 0; i < n; i++ {
		sum += candies[i]
	}

	return sum
}

// 思路2：山顶的糖果数就等于从左边的山底或右边的山底依次加 1