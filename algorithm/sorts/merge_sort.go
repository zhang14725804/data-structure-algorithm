package sorts

/*
	归并排序思想（空间换时间）：

	归并排序（Merge sort）是建立在归并操作上的一种有效的排序算法。该算法是采用分治法（Divide and Conquer）的一个非常典型的应用。

	分割：递归地把当前序列平均分割成两半。
	集成：在保持元素顺序的同时将上一步得到的子序列集成到一起（归并）。

	作为一种典型的分而治之思想的算法应用，归并排序的实现由两种方法：

	(1) 自上而下的递归（所有递归的方法都可以用迭代重写，所以就有了第 2 种方法）；
	(2) 自下而上的迭代；


	和选择排序一样，归并排序的性能不受输入数据的影响，但表现比选择排序好的多，因为始终都是 O(nlogn) 的时间复杂度。代价是需要额外的内存空间。

	(1) 自上而下的递归：
	申请空间，使其大小为两个已经排序序列之和，该空间用来存放合并后的序列
	设定两个指针，最初位置分别为两个已经排序序列的起始位置
	比较两个指针所指向的元素，选择相对小的元素放入到合并空间，并移动指针到下一位置
	重复步骤3直到某一指针到达序列尾
	将另一序列剩下的所有元素直接复制到合并序列尾
*/
func (sort Sort) MergeSort(array []int) []int{
	length:=len(array)
	if length<2{
		return array
	}
	key:=length/2
	left:=MergeSort(array[0:key])
	right:=MergeSort(array[key:])
	sort.result = merge(left,right)
}

func merge(left,right []int) []int {
	newArr:=make([]int,len(left)+len(right))
	i,j,index:=0,0,0
	for{
		if left[i]>right[j]{
			newArr[index] = right[j]
			index++
			j++
			if j == len(right){
				copy(newArr[index:],left[i:])
				break
			}
		}else{
			newArr[index] = left[i]
			index++
			i++
			if i == len(left){
				copy(newArr[index:],right[j:])
				break
			}
		}
	}
	return newArr
}
