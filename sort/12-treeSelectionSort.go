package main

import (
	"fmt"
	"math"
)

/*
	锦标赛排序，也称为树形选择排序（Tree Selection Sort）
	todo:复杂(不懂)
*/

type node struct{
	value int //叶子的数据
	isOk bool // 叶子状态是不是无穷大
	rank int // 叶子的排序
}

func main()  {
	var l = 10
	var hash = make(map[int]int,l)
	var obj []int
	// map内部随机存储
	for i:=0;i<l;i++{
		hash[i] = i
	}
	for k,_:=range hash{
		obj = append(obj,k)
	}
	fmt.Println(obj)
	fmt.Println(TreeSlectionSort(obj))
}

// 返回x的y次方
func pow(x,y int)int{
	return int(math.Pow(float64(x),float64((y))))
}

func TreeSlectionSort(arr []int) []int {
	var level int = 0// 树的层级
	var result = make([]int,0,len(arr)) // 保存最终结果
	for pow(2,level) < len(arr){
		level++ // 求出可以覆盖所有元素的层数
	}
	var leaf = pow(2,level) // 叶子数量
	var tree = make([]node,level*2-1) //构造树
	// 填充叶子（todo：leaf+i-1是什么意思） index out of range [15] with length 7 什么鬼
	for i:=0;i<len(arr);i++{
		tree[leaf+i-1] = node{value:arr[i],isOk:true,rank:i}
	}
	// 进行对比
	for i:=0;i<level;i++{
		nodeCount := pow(2,level-i) // 每次处理降低一个层级
		for j:=0;j<nodeCount/2;j++{
			leftNode := nodeCount-1+j*2
			rightNode := leftNode+1
			// 中间节点存储最小值
			if !tree[leftNode].isOk || (tree[rightNode].isOk && tree[leftNode].value>tree[rightNode].value){
				mid:=(leftNode-1)/2
				tree[mid].value = tree[rightNode].value
			}else{
				mid:=(leftNode-1)/2
				tree[mid].value = tree[leftNode].value
			}
		}

	}
	// 保存最顶端的最小数
	result = append(result,tree[0].value)
	for t:=0;t<len(arr)-1;t++{
		winnode := tree[0].rank+leaf-1 // 记录节点
		tree[winnode].isOk = false
		for i:=0;i<level;i++{
			leftNode := winnode
			if winnode%2 ==0{ // 处理奇偶数
				leftNode = winnode-1
			}
			rightNode := leftNode+1
			// 中间节点存储最小值
			if !tree[leftNode].isOk || (tree[rightNode].isOk && tree[leftNode].value>tree[rightNode].value){
				mid:=(leftNode-1)/2
				tree[mid].value = tree[rightNode].value
			}else{
				mid:=(leftNode-1)/2
				tree[mid].value = tree[leftNode].value
			}
			winnode = (leftNode-1)/2 // 保存中间节点
		}
		result = append(result,tree[0].value)
	}
	return result
}