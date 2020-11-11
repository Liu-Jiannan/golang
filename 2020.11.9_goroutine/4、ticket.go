package main

import (
	"fmt"
	"math/rand"
	"time"
)

//全局变量，表示票
var ticket = 100 //100张票

func main()  {
	/*
	4个goroutine,模拟4个售票口
	 */

	go saleTickets("窗口1")
	go saleTickets("窗口2")
	go saleTickets("窗口3")
	go saleTickets("窗口4")

	time.Sleep(5*time.Second)
}

func saleTickets(name string)  {
	for{
		rand.Seed(time.Now().UnixNano())
		if ticket >0 {
			time.Sleep(time.Duration(rand.Intn(100))*time.Millisecond)
			fmt.Println(name,"售出:",ticket)
			ticket--
		}else {
			fmt.Println(name,"售完，没有票了")
			break
		}
	}
}
