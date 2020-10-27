package main

import (
	"fmt"
	"io"
	"os"
)

func main()  {
	/*
	读取数据:
		Reader接口:
			Read(p []byte)(n int,error)
	 */
	//step1:打开文件
	fileName := "aa.txt"
	file,err:=os.Open(fileName)
	if err!=nil {
		fmt.Println("err:",err)
		return
	}
	//stap3:关闭文件
	defer file.Close()

	//step2:读取数据
	bs := make([]byte,4,4)
	/*
	//第一次读取
	n,err := file.Read(bs)
	fmt.Println(err)
	fmt.Println(n)
	fmt.Println(bs)
	fmt.Println(string(bs))

	//第二次读取
	n,err = file.Read(bs)
	fmt.Println(err)
	fmt.Println(n)
	fmt.Println(bs)
	fmt.Println(string(bs))

	//第三次打印
	n,err = file.Read(bs)
	fmt.Println(err)
	fmt.Println(n)
	fmt.Println(bs)
	fmt.Println(string(bs))
	*/
	n:= -1
	for{
		n,err = file.Read(bs)
		if n ==0 ||err == io.EOF{
			fmt.Println("读取到了文件的末尾，结束读取操作。。")
			break
		}
		fmt.Println(string(bs[:n]))
	}
}
