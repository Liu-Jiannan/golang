package main

import "fmt"


func main(){
	/*
		接口 interface
			在go中，接口是一组方法签名.

			当某个类型为这个接口中的所有方法提供了方法的实现，它被称为实现接口

			Go语言中, 接口与类型的实现关系，是非嵌入式

	1.当需要接口类型的对象时,可以使用任意实现类对象代替
	2.接口对象不能访问实现类中的属性
	*/

	//1. 创建Mouse类型
	m1 := Mouse{name:"黑色华硕"}
	fmt.Println(m1.name)
	//2. 创建FlashDisk类型
	f1 := FlashDisk{name:"华硕机械"}
	fmt.Println(f1.name)

	//调用测试方法
	testInterface(m1)
	testInterface(f1)

	//声明一个接口类型
	var usb USB
	usb= m1
	usb.start()
	usb.end()
	//usb不能打印字段

}



//1	定义接口
type USB interface {
	start()	//USB设备开始工作
	end()	//USB设备结束工作
}

//2	实现类
type Mouse struct {
	name string  //字段属性
}

type FlashDisk struct {
	name string
}

func (m Mouse)start(){
	fmt.Println(m.name,"鼠标,准备就绪,可以开始工作,滴滴滴...")
}

func (m Mouse)end(){
	fmt.Println(m.name,"鼠标结束工作,可以安全退出!")
}

func (f FlashDisk)start(){
	fmt.Println(f.name,"键盘,准备就绪,可以开始工作,哒哒哒...")
}

func (f FlashDisk)end(){
	fmt.Println(f.name,"键盘结束工作")
}

//3 测试方法
func testInterface(usb USB){
	usb.start()
	usb.end()
}