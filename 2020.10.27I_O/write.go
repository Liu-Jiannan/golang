package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		写出数据:
	*/
	fileName := "aa.txt"
	//打开文件
	//写出数据
	//关闭文件
	file, err := os.OpenFile(fileName,os.O_CREATE|os.O_WRONLY|os.O_APPEND,os.ModePerm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer file.Close()

	//bs := []byte{65, 66, 67, 68, 69, 70}
	bs :=  []byte{97,98,99,100}
	n, err := file.Write(bs[:2])
	fmt.Println(n)
	HandLeErr(err)

	//直接写出字符串
	n,err = file.WriteString("\nHelloWorld")
	fmt.Println(n)
	HandLeErr(err)

	//file.WriteString("\n")

	n,err=file.Write([]byte("\ntoday"))
	fmt.Println(n)
	HandLeErr(err)

}

func HandLeErr(err error)  {
	if err != nil {
		fmt.Println(err)
		return
	}
}
