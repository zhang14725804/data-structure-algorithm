package main

import "fmt"

// 构建大顶堆
func HeapSortMax(arr []int,l int) []int {
	if l <=1{
		return arr
	}else{
		depth := l/2 -1
		for i:=depth;i>=0;i--{
			topMax := i // 假定最大的在i位置
			leftChild := 2*i +1
			rightChild := 2*i + 2
			// 防止数组越界
			if leftChild<=l-1 && arr[leftChild] > arr[topMax]{
				topMax = leftChild
			}
			if rightChild <= l-1 && arr[rightChild] > arr[topMax]{
				topMax = rightChild
			}
			// 确保i的值最大
			if topMax != i{
				arr[i],arr[topMax] = arr[topMax],arr[i]
			}
		}
		return arr
	}
}

func HeapSort(arr []int) []int {
	l := len(arr)
	for i:=0;i<l;i++{
		// 每次截取一段
		last := l - i
		HeapSortMax(arr,last)
		// todo：这里略不懂（最大的数放在最后）
		if i<l{
			arr[0],arr[last-1] = arr[last-1],arr[0]
		}
	}
	return arr
}

func main()  {
	arr:=[]int{4,7,2,1,3,8,5,9,6,8}
	fmt.Println(HeapSort(arr))
}