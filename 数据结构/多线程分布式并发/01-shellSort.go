// 每个步长在进行插入排序
func step(arr []int, start, gap int) {
	l := len(arr)
	for i := start + gap; i < l; i += gap {
		// todo：不太懂
		backup := arr[i] // 备份插入的数据
		j := i - gap     // 上一个位置循环找到位置插入
		for j >= 0 && backup < arr[j] {
			arr[j+gap] = arr[j] // 从前往后移动
			j -= gap
		}
		// 插入
		arr[j+gap] = backup
	}
}

/*
	普通shell排序
*/
func ShellSort(arr []int) []int {
	l := len(arr)
	if l <= 1 {
		return arr
	} else {
		gap := l / 2
		for gap > 0 {
			//
			for i := 0; i < gap; i++ {
				step(arr, i, gap)
			}
			gap /= 2 // gap--
		}
	}
	return arr
}

/*
	多线程版的shell排序
*/
func ShellSortGoroutine(arr []int) []int{
	if len(arr)<2 || arr == nil{
		return arr
	}

	wg:=sync.WaitGroup{} // 等待多个线程返回
	grNum := runtime.NumCPU() // 系统有几个CPU

	for gap:=len(arr)/2;gap>0;gap/=2{
		wg.Add(grNum)
		ch := make(chan int,1000)
		go func(){
			// 管道写入任务
			for k:=0;k<gap;k++{
				ch <- k
			}
			close(ch) // 关闭管道
		}()
		for k:=0;k<grNum;k++{
			go func(){
				for v:=range ch{
					step(arr,v,gap) // 每一步长
				}
				wg.Done() // 
			}()
		}
		wg.Wait()
	}
	return arr
}

func main() {
	arr := []int{9, 2, 5, 0, 6, 1, 7, 8, 4, 3}
	fmt.Println(ShellSortGoroutine(arr))
}