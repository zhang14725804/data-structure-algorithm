/*
	给定一个非负整数数组 A，返回一个数组，在该数组中， A 的所有偶数元素之后跟着所有奇数元素。

	思路：双指针法，一个从头开始遍历，一个从尾部开始遍历，每次走一步，根据奇偶数交换位置或者向前向后移动指针
*/
func sortArrayByParity(arr []int) []int {
	i,j := 0,len(arr)-1
	// 边界条件
	for i<j{
		if arr[i]%2 == 1 && arr[j]%2 == 0{
			arr[i],arr[j] = arr[j],arr[i]
		}
		if arr[i] %2 == 0{
			i++
		}
		if arr[j] %2 == 1{
			j--
		}
	}
	return arr
}