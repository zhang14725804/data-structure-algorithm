/*
	给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
*/

/*
	解题思路（question 😅😅😅 ）：
	对于这个问题，我们可以先对集合求和，得出sum，把问题转化为背包问题：
	给一个可装载重量为sum/2的背包和N个物品，每个物品的重量为nums[i]。现在让你装物品，是否存在一种装法，能够恰好将背包装满？

	动态规划思路( question )：

	第一步要明确两点，「状态」和「选择」
	(1)状态就是「背包的容量」和「可选择的物品」
	(2)选择就是「装进背包」或者「不装进背包」

	第二步要明确dp数组的定义
	dp[i][j] = x表示，对于前i个物品，当前背包的容量为j时，若x为true，则说明可以恰好将背包装满，若x为false，则说明不能恰好将背包装满。

	第三步，根据「选择」，思考状态转移的逻辑
	(1)不把nums[i]算入子集:dp[i][j] = dp[i-1][j]
	(2)把nums[i]算入子集:dp[i][j] = dp[i - 1][j-nums[i-1]]
	由于i是从 1 开始的，而数组索引是从 0 开始的，所以第i个物品的重量应该是nums[i-1]
	dp[i - 1][j-nums[i-1]]也很好理解：你如果装了第i个物品，就要看背包的剩余重量j - nums[i-1]限制下是否能够被恰好装满
*/
func canPartition1(nums []int) bool {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	// 和为奇数时，不可能划分成两个和相等的集合
	if sum%2 != 0 {
		return false
	}
	sum = sum / 2
	// dp[i][j] = x表示，对于前i个物品，当前背包的容量为j时，若x为true，则说明可以恰好将背包装满，若x为false，则说明不能恰好将背包装满
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, sum+1)
	}
	// base case
	for i := 0; i <= n; i++ {
		dp[i][0] = true
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			// 背包容量不足，不能装入第 i 个物品
			if j-nums[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				// 装入或不装入背包
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[n][sum]
}

/*
	状态压缩1 ( 难懂😅😅😅😅😅😅)

	背包的体积为sum / 2
	背包要放入的商品（集合里的元素）重量为 元素的数值，价值也为元素的数值
	背包如何正好装满，说明找到了总和为 sum / 2 的子集。
	背包中每一个元素是不可重复放入。
*/
func canPartition(nums []int) bool {
    sum:=0
    for _,v:=range nums{
        sum+=v
    }
    if sum%2!=0{
        return false
    }
    sum = sum/2
	// 容量为j的背包，所背的物品价值可以最大为dp[j]
    dp:=make([]int,sum+1)
	// 如果使用一维dp数组，物品遍历的for循环放在外层，遍历背包的for循环放在内层，且内层for循环倒叙遍历！
    for i:=0;i<len(nums);i++{
        for j:=sum;j>=nums[i];j--{
			// 相当于背包里放入数值，那么物品i的重量是nums[i]，其价值也是nums[i]
            dp[j] = MaxInt(dp[j],dp[j-nums[i]]+nums[i])
        }
    }
    return dp[sum]==sum
}

/*
	question:状态压缩2(不懂 😅😅😅😅😅😅)
*/
func canPartition(nums []int) bool {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	// 和为奇数时，不可能划分成两个和相等的集合
	if sum%2 != 0 {
		return false
	}
	sum = sum / 2
	// dp[i]表示 背包总容量是i，最大可以凑成i的子集总和为dp[i]
	dp := make([]bool, sum+1)
	// base case
	dp[0] = true
	for i := 0; i < n; i++ {
		// j应该从后往前反向遍历，因为每个物品（或者说数字）只能用一次
		for j := sum; j >= 0; j-- {
			if j-nums[i] >= 0 {
				dp[j] = dp[j] || dp[j-nums[i]]
			}
		}
	}
	return dp[sum]
}