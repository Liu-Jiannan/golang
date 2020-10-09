package main

import "fmt"

type color struct {
	co string
}

func main()  {
//1 结构体作为返回值
	//结构体对象作为返回值
	f1 := getcolor1()
	f2 := getcolor1()
	fmt.Printf("更改前,f1,%v,%p\n,f2,%v,%p\n",f1,&f1,f2,&f2)
	f1.co="红"
	fmt.Printf("更改后,f1,%v,%p\n,f2,%v,%p\n",f1,&f1,f2,&f2)
	fmt.Println("---------------")

	//结构体指针作为返回值
	f3 := getcolor2()
	f4 := getcolor2()
	fmt.Printf("更改前f3:%v,%p\n,f4:%v,%p\n",f3,&f3,f4,&f4)
	f3.co="白"
	fmt.Printf("更改后f3:%v,%p\n,f4:%v,%p\n",f3,&f3,f4,&f4)
	fmt.Println("---------------")

//2 结构体作为参数
	f5 := color{co:"白"}
	fmt.Printf("更改前f5:%v,%p\n",f5,&f5)
	fmt.Println("---------------")
	//将结构体对象作为参数
	changeInfo1(f5)
	fmt.Printf("f5:%v,%p\n",f5,&f5)
	fmt.Println("-------------")
	//将结构体指针作为参数
	changeInfo2(&f5)
	fmt.Printf("f5:%v,%p\n",f5,&f5)
}

//返回结构体对象
func getcolor1() (f color) {
	f = color{co:"黄"}
	fmt.Printf("返回对象的函数内f1:%v,%p\n",f,&f)
	return
}

//返回结构体指针
func getcolor2() (f *color) {
	//f = &color{"黑"}
	temp := color{"黑"}
	fmt.Printf("函数内temp:%v,%p\n",temp,&temp)
	f = &temp
	fmt.Printf("返回函数内f:%v,%p\n",f,&f)
	return
}

func changeInfo1(f color)  {
	f.co = "蓝"
	fmt.Printf("函数changeInfo1内f:%v,%p\n",f,&f)
}

func changeInfo2(f *color)  {
	f.co="紫"
	fmt.Printf("函数changeInfo2内f:%v,%p\n",f,&f)
}