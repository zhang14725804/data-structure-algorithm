package main

import (
	"fmt"
	"math/rand"
)

/*
	todo：	快速排序的优化：快排涉及很多数组初始化、拼接操作
	快速排序结合二分查找进行搜索数据
	todo:解决快速排序吃内存问题
	快速排序改良版
*/
func mainQuickSortPlus()  {
	arr:=[]int{9,7,8,6,4,5,1,3,2,0}
	fmt.Println("未排序：",arr)
	quickSortPlus(arr)
	fmt.Println("排序后：",arr)
}
// 插入排序
func merge(arr []int,left,right int)  {

	for i:=left;i<=right;i++{
		temp:=arr[i] // 备份数据
		var j int
		for j=i;j>left && arr[j-1]>temp;j--{ // 定位
			arr[j] = arr[j-1] // 数据向后移动
		}
		arr[j] = temp //插入
	}
}
/*
	始终在一个数组交换元素位置（todo：不好理解），不用重新开辟连接数组
*/
func quickSortX(arr []int,left,right int)  {
	// 数组只剩下三个数，直接插入排序
	if right-left < 3 {
		merge(arr,left,right)
	}else{
		// 随机找一个数字
		swap(arr,left,rand.Int()%(right-left+1)+left)
		mid := arr[left] // 坐标数据，分成左右两部分
		lt := left // < mid
		gt := right+1 // > mid
		i := left+1 // == mid
		for i < gt {
			if arr[i] < mid{
				swap(arr,i,lt+1)
				lt++
				i++
			}else if arr[i] > mid{
				swap(arr,i,gt-1)
				gt--
			}else{
				i++
			}
		}
		swap(arr,left,lt) // 交换头部位置
		quickSortX(arr,left,lt-1) // 小于的部分
		quickSortX(arr,gt,right) // 大于的部分
	}

}

func quickSortPlus(arr []int)  {
	quickSortX(arr,0,len(arr)-1)
}

func swap(arr []int,i,j int)  {
	arr[i],arr[j] = arr[j],arr[i]
}