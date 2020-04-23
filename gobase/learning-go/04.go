func function(){
	/*
		函数作为值
	*/
	a:=func(){
		fmt.Println(0)
	}
	a()
	// 还能这么玩
	xs:=map[int]func()int{
		1:func()int{return 1},
		2:func()int{return 2},
		3:func()int{return 3},
	}
	fmt.Println(xs[1]())

	// 作为回调函数
	func callback(x int,cb func(int)){
		cb(x)
	}
	func printit(x int) {
		fmt.Printf("%v\n", x)
	}
	callback(3,printit)

	/*
		defer延迟函数（问题来了，defer什么原理）
	*/ 
	for i := 0; i < 5; i++ {
		// 0 1 2 3 4 4 3 2 1 0 
		fmt.Printf("%d ", i)
		defer fmt.Printf("%d ", i)
	}

	/* 可变参数 */
	func myFunc(args ...int){
		for _,arg := range args{
			fmt.Printf("the number is: %d\n", arg)
		}
	}
	myFunc(1,2,3)
}
