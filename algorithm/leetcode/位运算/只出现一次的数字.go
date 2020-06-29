/*
	给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

	思路：
	任意一个数和0异或仍然为自己
	任意一个数和自己异或是0
*/

func  singleNumber(nums []int) int {
	ans:=0
	for _,n:=range nums{
		ans^=n
	}
	return ans
}