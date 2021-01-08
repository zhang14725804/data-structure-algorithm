/*
	木桶排序（适用范围）
*/
func main() {
	arr := []int{1, 2, 3, 4, 4, 3, 2, 1, 3, 2, 4}
	fmt.Println(bucketSort(arr))
}

func bucketSort(arr []int) []int {
	l := len(arr)
	if l <= 1 {
		return arr
	}
	// 数据的范围
	num := 4
	// 二维数组
	buckets := make([][]int, num)
	for i := 0; i < l; i++ {
		// 木桶计数
		buckets[arr[i]-1] = append(buckets[arr[i]-1], arr[i])
	}
	fmt.Println(buckets)
	pose := 0
	for i := 0; i < num; i++ {
		bucketsLen := len(buckets[i])
		if bucketsLen > 0 {
			//
			copy(arr[pose:], buckets[i])
			//
			pose += bucketsLen
		}
	}
	return arr
}