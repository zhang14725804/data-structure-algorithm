package main

import "fmt"

/*
	插入排序（略微有点绕）
	从左往右，从右往左两种
*/
func InsertSort(arr []int) []int{
	l := len(arr)
	if l<=1{
		return arr
	}
	// 跳过第一个
	for i:=1;i<l;i++{
		backup := arr[i] // 备份插入的数据
		j:=i-1 // 上一个位置循环找到位置插入
		for j>=0 && backup<arr[j]{
			arr[j+1] = arr[j] // 从前往后移动
			j--
		}
		// 插入
		arr[j+1] = backup
	}
	return arr
}

func mainInsert()  {
	arr:=[]int{4,7,2,1,3,8,5,9,6,8}
	fmt.Println(InsertSort(arr))
}