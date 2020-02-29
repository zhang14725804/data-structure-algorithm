/*
	78. Subsets

	两种方法：
	（1）dfs方法
	（2）二进制思想（这个思路不错）
*/
func subsets(nums []int) [][]int {
	var res [][]int
	// 
	for i:=0;i<1<<len(nums);i++{
		var now []int
		for j:=0;j<len(nums);j++{
			// 位运算
			if i>>j & 1 == 1{
				now = append(now,nums[j])
			}
		}
		res = append(res,now)
	}
	return res
}