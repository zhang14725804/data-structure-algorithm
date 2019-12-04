package algorithm

import "fmt"

func SortMain() {
	fmt.Println("==========")
	origin := []int{5, 2, 7, 3, 6, 1}
	fmt.Println(origin)
	BubbleSort(origin)
}
