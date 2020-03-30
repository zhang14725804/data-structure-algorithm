package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
	内存搜索（更快）
*/

func mainMem()  {
	const N = 3
	// 初始化数组
	allStrs:=make([]string,N,N)
	i:=0// 统计多少行
	// 打开文件
	path:="./mock1.txt"
	file,_:=os.Open(path)
	defer file.Close()
	buf:=bufio.NewReader(file)
	// 数据载入内存
	for {
		// 逐行读取
		line,_,end:=buf.ReadLine()
		if end==io.EOF{
			break
		}
		allStrs[i] = string(line)
		i++
	}
	// todo：这里有问题，无法输入
	for ; ;  {
		fmt.Println("请输入要查询的数据")
		var keywords string
		// 获取键盘输入
		fmt.Scanln(keywords)
		// 内存搜索
		for j:=0;j<N ;j++  {
			if strings.Contains(allStrs[j],keywords){
				fmt.Println(allStrs[j])
			}
		}
	}

}
