package main

import "fmt"

/*
	使用接口定义汽车的标准：
	① 汽车品牌  ② 汽车能够驱动
	并定义结构体，实现卡车，电动汽车、三轮车三种定义，实现汽车接口标准。
	分别实例化解放牌卡车，特斯拉电动车，时风三轮车，并调用对应的驱动方法
	 */

func main()  {
	ka := kache{name:"解放牌卡车"}
	drivemethod(ka)

	dian := diandongqiche{"特斯拉电动车"}
	drivemethod(dian)

	san := sanlunche{name:"时风三轮车"}
	drivemethod(san)
}

//定义汽车标准
type car interface {
	brand() //品牌
	drive() //驱动
}

//实现汽车类别
type kache struct {
	name string
}

type diandongqiche struct {
	name string
}

type sanlunche struct {
	name string
}

//实现汽车接口标准
func (k kache)brand(){
	fmt.Println(k.name)
}

func (k kache)drive(){
	fmt.Println("卡车驱动正常")
}

func (d diandongqiche)brand()  {
	fmt.Println(d.name)
}

func (d diandongqiche)drive()  {
	fmt.Println("电动汽车驱动正常")
}

func (s sanlunche)brand(){
	fmt.Println(s.name)
}

func (s sanlunche)drive()  {
	fmt.Println("三轮车驱动正常")
}

//实现汽车驱动方法
func drivemethod(Car car)  {
	Car.brand()
	Car.drive()
}