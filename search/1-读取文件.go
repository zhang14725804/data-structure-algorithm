package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)
/*
	硬盘搜索（当前是硬盘搜索）
	内存搜索
*/
func mainDisk()  {
	// 打开文件
	path:="./mock1.txt"
	file,_:=os.Open(path)
	defer file.Close()
	buf:=bufio.NewReader(file)
	for {
		// 逐行读取
		line,_,end:=buf.ReadLine()
		if end==io.EOF{
			break
		}
		lineStr:=string(line)
		// 硬盘搜索
		if strings.Contains(lineStr,"88u"){
			fmt.Println(lineStr)
		}
	}
}
