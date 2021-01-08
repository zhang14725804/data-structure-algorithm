
/*
	休眠排序（todo：一脸懵逼）
*/
var flag bool
var container chan bool
var count int

func mainSleepSort() {
	array := []int{16, 8, 1, 24, 30}
	flag = true
	// 管道
	container = make(chan bool, 5)

	for i := 0; i < len(array); i++ {
		go sleeping(array[i])
	}

	go listen(len(array))
	for flag {
		time.Sleep(time.Second)
	}
}
func listen(size int) {
	for flag {
		select {
		case <-container:
			count++
			// 数据采集完成，退出
			if count >= size {
				flag = false
				break
			}
		}
	}
}

// 注意参数类型
func sleeping(data int) {
	time.Sleep(time.Duration(data) * time.Microsecond * 1000)
	fmt.Println(data)
	container <- true
}