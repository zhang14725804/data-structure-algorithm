/*
	来源：https://www.youtube.com/watch?v=LvgVSSpwND8
*/
func main()  {
	/*
		count("sleep")
		count("fish")
		// 顺序执行
		0 sleep
		1 sleep
		2 sleep
		0 fish
		1 fish
		2 fish
	*/
	/*
		go count("sleep")
		count("fish")
		// 交替执行
		0 fish
		0 sleep
		1 sleep
		1 fish
		2 sleep
		2 fish
	*/
	/*
		go count("sleep")
		go count("fish")
		// 都不执行（todo：为什么）
	*/
	/*
		go count("sleep")
		go count("fish")
		time.Sleep(time.Second)
		// 执行一部分
		0 fish
		0 sleep
		1 fish
		1 sleep
	*/
	/*
		go count("sleep")
		go count("fish")
		fmt.Scanln()
		// 按下回车之前执行程序，之后中断
	*/

	/*
		// 只执行一次
		c:=make(chan string)
		go count2("sleep",c)
		msg:=<-c
		fmt.Println(msg)
	*/
	/*
		// 执行完所有
		c:=make(chan string)
		go count2("sleep",c)
		for{
			msg,open:=<-c
			// 解决CPU狂飙问题
			if !open{
				break
			}
			fmt.Println(msg)
		}

		// 执行完所有
		c:=make(chan string)
		go count2("sleep",c)
		// 解决goroutine关闭问题
		for msg:= range c{
			fmt.Println(msg)
		}
	*/
	/*
		// fatal error: all goroutines are asleep - deadlock!
		c := make(chan string)
		c <-"hello"
		msg:=<-c
		fmt.Println(msg)

		// todo：为什么这样就可以执行了
		c := make(chan string,1)
		c <-"hello"
		msg:=<-c
		fmt.Println(msg)

		c := make(chan string,2)
		c <-"hello"
		c <-"world"

		msg:=<-c
		fmt.Println(msg)

		msg=<-c
		fmt.Println(msg)
		// todo:fatal error: all goroutines are asleep - deadlock!
		c := make(chan string,2)
		c <-"hello"
		c <-"world"
		c <- "mino"
		msg:=<-c
		fmt.Println(msg)

		msg=<-c
		fmt.Println(msg)
	*/
	/*
	c1:=make(chan string)
	c2:=make(chan string)

	go func(){
		for{
			c1<-"每隔 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func(){
		for{
			c2<-"每隔 2000ms"
			time.Sleep(time.Millisecond * 2000)
		}
	}()

	for{
		// 间隔输出
		//fmt.Println(<-c1)
		//fmt.Println(<-c2)
		// 根据时间间隔输出
		select {
			case msg1:=<-c1:
				fmt.Println(msg1)
			case msg2:=<-c2:
				fmt.Println(msg2)
		}
	}
	 */
	jobs := make(chan int,100)
	result := make(chan int, 100)
	// todo：goroutine威力：协程
	go worker(jobs, result)
	go worker(jobs, result)
	go worker(jobs, result)
	go worker(jobs, result)
	for i:=0;i<100 ;i++  {
		jobs <- i
	}
	close(jobs)
	for i:=0;i<100 ; i++ {
		fmt.Println(<-result)
	}
}
func fib(n int) int {
	if n<=1{
		return n
	}
	return fib(n-1)+fib(n-2)
}
func worker(jobs <- chan int, result chan int)  {
	for n := range jobs{
		// Go中channel可以是只读、只写、同时可读写的 (参数:result <- chan int报错)invalid operation: result <- fib(n) (send to receive-only type <-chan int)
		result <- fib(n)
	}
}
func count(thing string)  {
	for i:=0;i<5 ;i++  {
		fmt.Println(i,thing)
		time.Sleep(time.Millisecond*500)
	}
}
func count2(thing string,c chan string)  {
	for i:=0;i<5 ;i++  {
		c<-thing
		time.Sleep(time.Millisecond*500)
	}
	// fatal error: all goroutines are asleep - deadlock!（解决这个问题，但是CPU狂飙）
	close(c)
}