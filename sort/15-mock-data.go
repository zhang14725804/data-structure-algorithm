// mock 生成大量数据
func main() {
	const LEN = 10000000
	// 新建文件
	// path:="./mock-big-data.txt"
	// saveFile,_:=os.Create(path)
	// defer saveFile.Close()
	// save:=bufio.NewWriter(saveFile)
	// for i:=0;i<LEN;i++{
	// 	// 写入文件
	// 	fmt.Fprintln(save,rand.Intn(LEN))
	// }
	// save.Flush() // 刷新

	target := make([]int, LEN, LEN)
	for i := 0; i < LEN; i++ {
		target[i] = rand.Intn(LEN)
	}
	// 统计计算时间
	t1 := time.Now()
	// 快速排序时间 3.8417336s
	sortedTarget := QuickSortMock(target)
	// 插入排序特别慢
	used := time.Since(t1)
	fmt.Println(used)

	// 新建文件并写入
	path := "./mock-big-data-sorted.txt"
	saveFile, _ := os.Create(path)
	defer saveFile.Close()
	save := bufio.NewWriter(saveFile)
	for i := 0; i < LEN; i++ {
		// 写入文件
		fmt.Fprintln(save, sortedTarget[i])
	}
	save.Flush() // 刷新
}
func QuickSortMock(arr []int) []int {
	l := len(arr)
	if l <= 1 {
		return arr
	} else {
		// base可以取随机一个数，也可以随机生成一个数
		base := arr[0]
		smaller := make([]int, 0, 0)
		bigger := make([]int, 0, 0)
		mid := make([]int, 0, 0)
		mid = append(mid, base)
		for i := 1; i < l; i++ {
			if arr[i] < base {
				smaller = append(smaller, arr[i])
			} else if arr[i] > base {
				bigger = append(bigger, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		// 递归处理(todo:这里可以用栈实现)
		smaller, bigger = QuickSortMock(smaller), QuickSortMock(bigger)
		array := append(append(smaller, mid...), bigger...)
		return array
	}
}