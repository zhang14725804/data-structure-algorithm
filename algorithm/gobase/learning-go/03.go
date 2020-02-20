func arraySliceMap(){
	/*
		数组(不指定长度，默认len为0)
		var arr []int   // runtime error: index out of range [0] with length 0
	*/ 
	var arr [3]int
	arr[0] = 0
	arr[1] = 1
	fmt.Println(arr)

	/*
		数组类型具有固定的大小。大小是类型的一部分。它们无法生长，因为那样它们将具有不同的类型。
		数组也是值：将一个数组分配给另一个数组将复制所有元素。特别是，如果将数组传递给函数，它将接收该数组的副本，而不是指向它的指针

		 The size is part of the type. They can’t grow, because then they would have a different type. Also arrays are values: Assigning one array to another copies all the elements.
		  In particular, if you pass an array to a function it will receive a copy of the array, not a pointer to it.
	*/

	// 声明数组
	var arr2 [2]int
	arr3:=[1]int{1}
	arr4:=[...]int{1,2,3}
	fmt.Println(arr2,arr3,arr4)

	/*
		切片（slice）
		A slice is similar to an array, but it can grow when new elements are added. A slice always refers to an underlying array. 
		What makes slices different from arrays is that a slice is a pointer to an array; slices are reference types.
		如果您将一个切片分配给另一个切片，则它们都引用同一个 基础数组
	*/
	arr5 := [...]int{1,2,3,4,5,6,7,8} // 数组
	slice1:=arr5[:8]
	fmt.Println(len(slice1),cap(slice1))

	// cap和len的区别
	slice2 := []int{2, 3, 5, 7, 11, 13} // slice
	printSlice(slice2)

	// Slice the slice to give it zero length.
	slice2 = slice2[:0]
	printSlice(slice2)

	// Extend its length.
	slice2 = slice2[:4]
	printSlice(slice2)

	// Drop its first two values.
	slice2 = slice2[2:]
	printSlice(slice2)

	/*
		copy :: 复制功能将切片元素从源复制到目标，并返回其复制的元素数。此数字是源长度和目标长度的最小值
	*/
	var ac = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	var sc = make([]int, 6)
	n1 := copy(sc, ac[0:]) 
	n2 := copy(sc, sc[2:]) 
	// 6,4
	fmt.Println(n1,n2)

	/*
		map没什么
	*/
}