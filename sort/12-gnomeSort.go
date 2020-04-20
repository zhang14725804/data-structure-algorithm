package main

import "fmt"

/*
	侏儒排序：冒泡+插入
*/
func main()  {
	arr:=[]int{9,7,8,6,4,5,1,3,2,0}
	fmt.Println(gnomeSort(arr))
}

func gnomeSort(arr []int) []int  {
	i:=1
	for ;i<len(arr);{
		if arr[i]>=arr[i-1]{
			i++ // 符合顺序，继续前进
		}else{
			arr[i],arr[i-1] = arr[i-1],arr[i]
			//
			if i>1{
				i--
			}
		}
	}
	return arr
}