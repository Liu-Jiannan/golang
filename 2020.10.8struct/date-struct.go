package main

import "fmt"

	//定义Teacher结构体
type Teacher struct {
	name string
	age int8
	sex byte
}

func main()  {

	// 1. var声明方式实例化结构体,初始化方式为:		对象.属性=值
	var t1 Teacher
	fmt.Println(t1)
	//	%T 打印字符串		%v 打印整型		%q 打印字节
	fmt.Printf("t1:%T,%v,%q\n",t1,t1,t1)
	t1.name = "剑南春"
	t1.age = 18
	t1.sex = 1
	fmt.Println(t1)
	fmt.Printf("t1:%T,%v,%q\n",t1,t1,t1)
	fmt.Println("--------------")

	// 2. 变量简短声明格式实例化结构体,初始化方式: 对象.属性=值
	t2 := Teacher{}
	t2.name = "David"
	t2.age = 30
	t2.sex = 1
	fmt.Println(t2)

	// 3. 变量简短声明格式实例化结构体,声明时初始化,初始化方式为: 属性:值,属性:值  可以同行，也可以换行(类似map的用法)
	t3 := Teacher{
		name: "Josh",
		age:  31,
		sex:  1,
	}
	fmt.Println(t3)

	// 4. 变量简短声明格式实例化结构体,声明时初始化,不写属性名,按属性顺序只写属性值
	t4 := Teacher{"Ruby",30,0}
	fmt.Println(t4)
}
