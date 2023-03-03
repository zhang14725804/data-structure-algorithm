
/*
	方法一：枚举
	1. 排序
	2. begin 指向第一个连续相同元素的子序列的第一个元素， end 指向相邻的第二个连续相同元素的子序列的末尾元素
	   如果满足二者的元素之差为 1 ，则当前的和谐子序列的长度即为两个子序列的长度之和
	   等于  end−begin+1
*/
func findLHS(nums []int) (ans int) {
	sort.Ints(nums)
	begin := 0
	for end, num := range nums {
		for num-nums[begin] > 1 {
			begin++
		}
		if num-nums[begin] == 1 && end-begin+1 > ans {
			ans = end - begin + 1
		}
	}
	return
}

/*
	方法一：枚举
	方法二：哈希表
	1. map记录元素出现的次数
	2. 遍历map，假设(key,value),统计value+1出现的次数，count(value)+count(value+1)
*/
func findLHS(nums []int) int {
	cnt := map[int]int{}
	for _, n := range nums {
		cnt[n]++
	}

	ans := 0
	for num, ct := range cnt {
		if c1 := cnt[num+1]; c1 > 0 && c1+ct > ans {
			ans = c1 + ct
		}
	}
	return ans
}