/*
	给定一个非负整数数组 A， A 中一半整数是奇数，一半整数是偶数。
	对数组进行排序，以便当 A[i] 为奇数时，i 也是奇数；当 A[i] 为偶数时， i 也是偶数。

	思路：双指针，两个指针，每次移动两位，根据奇偶数交换位置或者向后移动指针
*/
func sortArrayByParityII(arr []int) []int {
	i,j := 0,1
	// 边界条件
	for ;i <len(arr) ;i+=2{
		// 如果i位是奇数
		if arr[i] % 2 == 1{
			// // 如果j位是奇数，移动j位，直到找到奇数位为止
			for arr[j] % 2 == 1 && j<len(arr){
				j+=2
			}
			arr[i],arr[j] = arr[j],arr[i]
		}
	}
	return arr
}