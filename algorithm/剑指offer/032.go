/*
	调整数组顺序使奇数位于偶数前面
	todo：双指针算法（一个奇数，一个偶数）
	我的想法：一个for循环，两个数组，一个存奇数，一个存偶数，循环完了拼接数组（low）
*/
func reOrderArray(array []int){
	i,j := 0,len(array)-1
	for i<=j{
		for i<=j && array[i] % 2 == 1{
			i++
		}
		for i<=j && array[j] % 2 == 0{
			j--
		}
		if i<j{
			swap(array,i,j)
		}
	}
}
func swap(arr []int,x,y int){
	arr[x],arr[y] = arr[y],arr[x]
}