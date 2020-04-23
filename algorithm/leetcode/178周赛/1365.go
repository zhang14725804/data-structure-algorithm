/*
	1365
	我的思路：遍历数组中每个数字，对于每个数字再遍历数组，判断是否比当前的数字小，放入结果数组中
*/

func smallerNumbersThanCurrent(nums []int) []int {
	n := len(nums)
	res:=make([]int,n)
	for i := 0; i < n; i++ {
		for _,num:=range nums{
			if nums[i]>num{
				res[i]++
			}
		}
	}
	return res
}