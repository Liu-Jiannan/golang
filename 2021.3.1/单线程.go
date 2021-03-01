package main

import (
	"fmt"
	"time"
)

//主线程 main线程 (单线程)
func main()  {
	/**
		计算机:任务。计算要实现的工作.
		任务管理器:1、goland 2typora 3、wps office
		windows，mac，linux:多任务操作系统
			进程:计算机管理应用的调度单位
			线程:再进程当中，至少应该有一个主线程.

		具体:打开goland软件
		1.新建一个进程，给进程一个编号。
		2.在进程中有一个主线程:运行程序功能
		3.下载功能：后台默默执行.(开启一个新的子线程)

		程序顺序:顺序执行(从上到小,从前到后)
			1、跑步 系鞋带 继续跑: 并发指的是多个事件在同一时间间隔发生
			2、跑步+听音乐:并行指的是多个事件在同一时刻发生


		单核cpu、多核cpu(8核\16核)

	*/

	//1、demo02.go  源代码
	//2、编译:gobuild   demo02.go --> demo02.exe
	//3、.exe : qq、wechat -->运行 -->启动进程
	//4、进程当中有一个主线程:main函数所执行的线程
	//顺序结构执行
	fmt.Println("hello...")
	go test1()
	go hellowold()
	time.Sleep(500)


}

func test1()  {
	fmt.Println("test被调用")

}

func hellowold()  {
	fmt.Println("hellowold执行")
}