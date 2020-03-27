package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main()  {
	// 打开文件
	file,err:=os.Open("./readWrite.sql")
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer file.Close()

	// 新建文件
	path:="./readWrite.txt"
	saveFile,_:=os.Create(path)
	defer saveFile.Close()
	save:=bufio.NewWriter(saveFile)

	// 读取文件流
	buf :=bufio.NewReader(file)
	for{
		line,_,err:=buf.ReadLine()
		if err==io.EOF{
			break
		}
		lineStr := string(line)
		strs:=strings.Split(lineStr," # ")
		fmt.Println(strs[1])
		// 写入文件
		fmt.Fprintln(save,strs[1])
	}
	save.Flush() // 刷新
}