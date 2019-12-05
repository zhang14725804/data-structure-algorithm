package sorts

import "fmt"

/*
	希尔排序思想(TODO：不理解)：


*/
func ShellSort(array []int) {
	length := len(array)
	if length < 2 {
		return
	}
	key := length / 2
	for key > 0 {
		for i := key; i < length; i++ {
			j := i
			for j >= key && array[i] < array[j-key] {
				array[i], array[j-key] = array[j-key], array[i]
				j = j - key
			}
		}
		key = key / 2
	}
	fmt.Println("希尔排序之后：：", array)
}
