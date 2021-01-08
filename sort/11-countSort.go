 

import "fmt"

func selectMaxy(arr []int) int {
	l := len(arr)
	if l <= 1 {
		return arr[0]
	} else {
		max := arr[0]
		for i := 1; i < l; i++ {
			if arr[i] > max {
				max = arr[i]
			}
		}
		return max
	}
}

/*
	统计排序（todo:代码不理解）
*/
func CountSort(arr []int) []int {
	max := selectMaxy(arr)
	sorted := make([]int, len(arr))
	counts := make([]int, len(arr))

	for _, v := range arr {
		counts[v]++
	}
	// 这几个意思
	for i := 1; i <= max; i++ {
		counts[i] += counts[i-1]
	}

	for _, v := range arr {
		sorted[counts[v]-1] = v
		counts[v]--
	}
	return sorted
}
func main() {
	arr := []int{1, 2, 3, 4, 5, 4, 3, 2, 1, 2, 3, 4, 5, 4, 3, 2, 1, 2}
	fmt.Println(CountSort(arr))
}