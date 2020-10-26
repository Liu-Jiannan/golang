package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	//1.时间戳:指定的日期，距离1970年1月1日0点0时0分0秒的时间差值：秒,纳秒
	t1 := time.Date(1970,1,1,1,0,0,0,time.UTC)
	timeStamp1 := t1.Unix()//秒的差值
	fmt.Println(timeStamp1)
	t0 := time.Now()
	timeStamp2 := t0.Unix()
	fmt.Println(timeStamp2)

	timeStamp3 := t1.UnixNano() //纳秒
	fmt.Println(timeStamp3)

	//2.时间间隔
	t2 :=t0.Add(time.Minute)//比现在时间多1分钟
	t3 :=t0.Add(24*time.Hour)//比现在时间多1天
	fmt.Println(t2)
	fmt.Println(t3)
	t4 := t0.AddDate(1,0,0)//添加指定的间隔
	fmt.Println(t4)

	d1:=t2.Sub(t0)
	fmt.Println(d1)//t0与t2的时间间隔

	//3.睡眠
	time.Sleep(3 *time.Second)//让当前的程序进入睡眠状态
	fmt.Println("main。。。over。。。")

	//睡眠(1-10)的随机秒数
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(10)+1
	time.Sleep(time.Duration(randNum)*time.Second)
}