 

import (
	"fmt"
	"strings"
)

/*
	golang中直接比较字符串，是比较内存地址，准确性不定（go1.3）
	strings.Compare()相对靠谱一些
*/
func SelectSortString(arr []string) []string {
	l := len(arr)
	if l <= 1 {
		return arr
	}

	// 只剩下第一个元素，不需要挑选
	for i := 0; i < l-1; i++ {
		min := i // 标记索引
		// 每次选出一个极小值
		for j := i + 1; j < l; j++ {
			// 同样的效果
			// if arr[min]<arr[j]{
			if strings.Compare(arr[min], arr[j]) > 0 {
				min = j // 保存极小值索引
			}
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}
func mainString() {
	arr := []string{"x", "m", "r", "a", "l", "y", "d"}
	fmt.Println(SelectSortString(arr))
}