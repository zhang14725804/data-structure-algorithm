package sorts

import "fmt"

/*
	参考维基百科：https://zh.wikipedia.org/wiki/%E5%B8%8C%E5%B0%94%E6%8E%92%E5%BA%8F

	插入排序在对几乎已经排好序的数据操作时，效率高，即可以达到线性排序的效率
	但插入排序一般来说是低效的，因为插入排序每次只能将数据移动一位

	分区域进行插入，就是希尔排序。采用分治法（Divide and Conquer）

	步长的选择是希尔排序的重要部分。只要最终步长为1任何步长序列都可以工作。算法最开始以一定的步长进行排序。
	然后会继续以一定步长进行排序，最终算法以步长为1进行排序。当步长为1时，算法变为普通插入排序，这就保证了数据一定会被排序

	原始数据：[ 13 14 94 33 82 25 59 94 65 23 45 27 73 25 39 10 ]
	第一次，以步长为5开始进行排序：
		13 14 94 33 82
		25 59 94 65 23
		45 27 73 25 39
		10
	对每列进行排序，排序后：
		10 14 73 25 23
		13 27 94 33 39
		25 59 94 65 82
		45
	第二次，以3为步长进行排序：
		10 14 73
		25 23 13
		27 94 33
		39 25 59
		94 65 82
		45
	对每列进行排序，排序后：
		10 14 13
		25 23 33
		27 25 59
		39 65 73
		45 94 82
		94
	最后以1步长进行排序（此时就是简单的插入排序了）
*/
func (sort Sort) ShellSort(array []int) {
	length := len(array)
	if length < 2 {
		return
	}
	// 步长
	key := length / 2
	for key > 0 {
		for i := key; i < length; i++ {
			j := i
			// 对每一列进行排序
			for j >= key && array[i] < array[j-key] {
				array[i], array[j-key] = array[j-key], array[i]
				j = j - key
			}
		}
		key = key / 2
	}
	sort.result = array
	// fmt.Println("希尔排序之后：：", array)
}
