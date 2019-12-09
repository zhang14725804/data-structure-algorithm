package sorts

/*
	快速排序思想：：

	首先要找一个数字作为基准数（这只是个专用名词）。
	为了方便，我们一般选择第 1 个数字作为基准数（其实选择第几个并没有关系）。
	接下来我们需要把这个待排序的数列中小于基准数的元素移动到待排序的数列的左边，把大于基准数的元素移动到待排序的数列的右边。
	这时，左右两个分区的元素就相对有序了；
	接着把两个分区的元素分别按照上面两种方法继续对每个分区找出基准数，然后移动，直到各个分区只有一个数时为止
*/
func QuickSort2(array []int) []int {
	if len(array)<2{
		return array
	}

	// 基准元素
	pivot:=array[0]
	var small,large,result []int
	for _,value :=range array[1:]{
		if value<=pivot{
			small=append(small,value)
		}else{
			large=append(large,value)
		}
	}
	result = append(result,QuickSort2(small)...)
	result = append(result,pivot)
	result = append(result,QuickSort2(large)...)
	return result
}
