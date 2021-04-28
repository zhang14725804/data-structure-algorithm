/*
	线程安全问题（多个线程访问一个资源的问题）
*/

var money int = 0

func add(val *int)  {
	for i := 0; i < 1000; i++ {
		*val++
	}
}

func main(){
	for i := 0; i < 1000; i++ {
		// 并发执行
		go add(&money)
	}

	fmt.Println(money)
}