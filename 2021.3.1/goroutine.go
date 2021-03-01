package main

import (
	"fmt"
	"time"
)

func main()  {
	/**
		goroutine:并发 go语言并发机制:协程，即goroutine
		进程、线程、协程
		进程:操作系统任务调度的基本单位.
		线程:线程存在于进程当中，一个进程当中可以有多个线程
		协程:细微级别的线程，一个线程当中可以有多个协程,
		用法:go 函数
	 */

	/**
	协程打印1000个数字,另外一条协程打印1000个字母.
	 */
	go printNum()
	go printCode()
	time.Sleep(10*time.Second)
}

func printNum() {
	for index := 0; index <1000;index++{
		fmt.Println("数字%d\n",index)
	}
}

func printCode()  {
	//字符编码:97 - 122
	codeArr := [6]string{"a","b","c","d","e","f"}
	for index := 0; index <1000 ; index++ {
		fmt.Println("字母:%s\n",codeArr[index%6])
	}
}
