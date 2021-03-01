package main

import (
	"fmt"
	"time"
)

func main()  {
	// 程序在main执行过程当中,我想下载一个文件
	/**
	并发编程:
		goroutine
	 */
	fmt.Println("hello world")
	//go是一个关键字
	go hello() //goroutine

	//保证子线程能执行,让主线程等一等
	time.Sleep(1000)//睡了1s

}

func hello()  {
	fmt.Println("hello,goroutine")
}