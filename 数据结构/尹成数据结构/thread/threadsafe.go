/*
	线程安全问题（多个线程访问一个资源的问题）
*/

var money int = 0
// 
var lock sync.Mutex

func add(val *int)  {
	lock.Lock()
	for i := 0; i < 1000; i++ {
		*val++
	}
	lock.Unlock()
}

func main(){
	for i := 0; i < 1000; i++ {
		// 并发执行
		go add(&money)
	}
	// 不加sleep结果不对，todo
	time.Sleep(time.Second*4)
	fmt.Println(money)
}