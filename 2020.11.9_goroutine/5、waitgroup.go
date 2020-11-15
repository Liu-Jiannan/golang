package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
func main() {
	/*
	WaitGroup:同步等待组
		Add(),设置等待组中要执行的子goroutine的数量

		Wait(),让主goroutine出于等待
	 */

	wg.Add(2)

	go fun1()
	go fun2()

	fmt.Println("main进入阻塞状态。。等待wg中的子goroutine结束。。")
	wg.Wait() //表示main goroutine进入等待，意味着阻塞
	fmt.Println("main解除阻塞")
}

func fun1()  {
	for i:= 1;i<10;i++{
		fmt.Println("fun1()函数中打印。。A",i)
	}
	wg.Done() //给wg等待组中的counter数值减1.

}

func fun2()  {
	defer wg.Done()
	for j:=1;j<10 ; j++ {
		fmt.Println("fun2()函数中打印。。B",j)
	}
}