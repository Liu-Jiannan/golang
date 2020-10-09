package main

import "fmt"

type Dog struct {
	name string
}

func main()  {
	//1. 实现结构体的深拷贝
	//struct是值类型,默认的复制就是深拷贝
	d1 := Dog{name:"大黄"}
	fmt.Printf("d1: %v,%p\n",d1,&d1)
	//深拷贝
	d2 := d1
	fmt.Printf("d2: %v,%p\n",d2,&d2)
	//d2修改值
	d2.name="小黑"
	fmt.Printf("d2: %v,%p\n",d2,&d2)

	//2.实现结构体浅拷贝:直接赋值指针地址
	d3 := &d1
	fmt.Printf("d3: %v,%p\n",d3,&d3)
	d3.name= "小白"
	fmt.Printf("d1: %v,%p\n",d1,&d1)
	fmt.Println("-------------------------")
	//3.实现结构体浅拷贝: 通过new()函数来实例化对象
	d4 := new(Dog)
	d4.name = "小花"
	d5 := d4
	fmt.Printf("d4: %v,%p\n",d4,&d4)
	fmt.Printf("d5: %v,%p\n",d5,&d5)
	fmt.Println("--------------")
	d5.name = "大胖"
	fmt.Println("d5修改后",d5)
	fmt.Println("d4",d4)
}
