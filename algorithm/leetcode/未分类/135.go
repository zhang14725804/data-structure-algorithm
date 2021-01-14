/*
	老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。

	你需要按照以下要求，帮助老师给这些孩子分发糖果：

	每个孩子至少分配到 1 个糖果。
	相邻的孩子中，评分高的孩子必须获得更多的糖果。

	那么这样下来，老师至少需要准备多少颗糖果呢？

	如果当前小朋友的 rating 比后一个小朋友的大，那么理论上当前小朋友的糖要比后一个的小朋友的糖多，但此时后一个小朋友的糖还没有确定，怎么办呢？

*/

// 思路1：利用正着遍历，再倒着遍历的思想
func candy(ratings []int) int {
	n := len(ratings)
	candies := make([]int, n)

	// 没人先发一个糖
	for i := 0; i < n; i++ {
		candies[i] = 1
	}

	// 正着进行
	for i := 0; i < n-1; i++ {
		if ratings[i] < ratings[i+1] {
			candies[i+1] = candies[i] + 1
		}
	}

	// 倒着进行
	for i := n - 1; i > 0; i-- {
		if ratings[i] < ratings[i-1] {
			// 后一个小朋友的糖果没有前一个的多，就更新后一个等于前一个加 1
			if candies[i-1] <= candies[i] {
				candies[i-1] = candies[i] + 1
			}
		}
	}

	sum := 0
	for i := 0; i < n; i++ {
		sum += candies[i]
	}

	return sum
}

// 思路2：山顶的糖果数就等于从左边的山底或右边的山底依次加 1