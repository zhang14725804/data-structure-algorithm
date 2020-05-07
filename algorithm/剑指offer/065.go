/*
	数组中的逆序对
	在数组中的两个数字如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。
	思路1：两个for循环暴力枚举
	todo：思路2：归并排序的思想，复杂呀😅。代码有问题
*/
func inversePairs(nums []int) int {
	res:=0
	for i:=0;i<len(nums);i++{
		for j:=i+1;j<len(nums);j++{
			if nums[i]>nums[j]{
				res++
			}
		}
	}
	return res
}


func merge(nums []int, l,r int) int{
	if l>=r{
		return 0
	}
	//
	mid := (l+r) >> 1
	res := merge(nums, l, mid) + merge(nums, mid+1, r)
	i := l
	j := mid+1
	var temp []int
	for i <= mid && j <= r{
		if nums[i]<= nums[j]{
		    i++
			temp = append(temp, nums[i])
		}else{
		    j++
			temp = append(temp, nums[j])
			res += mid-i+1
		}
	}
	for i <= mid{
	    i++
		temp = append(temp, nums[i])
	}
	for j <= r{
	    j++
		temp = append(temp, nums[j])
	}
	i = l
	for _,x := range temp{
	    i++
		nums[i] = x
	}
	return res
}
func inversePairs(nums []int) int {
	return merge(nums, 0, len(nums)-1)
}