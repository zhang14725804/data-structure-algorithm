package main

import "fmt"

/*
	鸡尾酒排序(双向冒泡排序)
	正向冒泡反向冒泡，每次得到极大值极小值
*/
func CocktailSort(arr []int) []int  {
	// 每次循环正向反向各冒泡一次
	for i:=0;i<len(arr)/2;i++{
		// 左右边界
		left:=0
		right:=len(arr)-1
		// 循环结束的条件
		for left<=right{
			if arr[left]>arr[left+1]{
				arr[left],arr[left+1] = arr[left+1],arr[left]
			}
			left++
			if arr[right-1]>arr[right]{
				arr[right-1],arr[right] = arr[right],arr[right-1]
			}
			right--
		}
		fmt.Println(i,arr)
	}
	return arr
}
func main()  {
	arr:=[]int{9,2,5,0,6,1,7,8,4,3}
	fmt.Println(CocktailSort(arr))
}
