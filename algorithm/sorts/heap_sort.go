package sorts

import "fmt"

/*
	堆排序思想(todos：不好理解)：

	堆分最大堆、最小堆。 以最大堆为例。

	最大堆是一个完全二叉树。 并且要求每个结点的值必须大于他的两个子节点。 所以他的根结点一定是最大值。

	但是左右结点大小不一定。

	用数组表示的二叉树，可以这样表达： i的子节点下标为 2*i + 1 和 2 * i + 2.   i的父节点下标为 (i-1)/2。
	对于数组长度length，大于length/2的下标一定没有子节点.

	排序思想是构建最大堆，之后根节点就是最大的一个了，把根结点拿出来，再把剩下的堆整理成最大堆，再把根拿出来。循环直到最后一个元素。
*/
func HeapSort(arr []int) {
	arrLen := len(arr)
	buildMaxHeap(arr, arrLen)
	for i := arrLen - 1; i >= 0; i-- {
		swap(arr, 0, i)
		arrLen -= 1
		heapify(arr, 0, arrLen)
	}
	fmt.Println("堆排序之后：：", arr)
}

func buildMaxHeap(arr []int, arrLen int) {
	for i := arrLen / 2; i >= 0; i-- {
		heapify(arr, i, arrLen)
	}
}

func heapify(arr []int, i, arrLen int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < arrLen && arr[left] > arr[largest] {
		largest = left
	}
	if right < arrLen && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		swap(arr, i, largest)
		heapify(arr, largest, arrLen)
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
