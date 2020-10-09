package main

import "fmt"

type student struct {
	name string
	age int8
}

func main()  {
	a := student{name:"Liu",age:19}
	a.nominalage()

}

//方法:方法名可以相同，但是接收的类型不能一样
func (a student) nominalage(){
	a.age=a.age+1
	fmt.Println(a)
}
//函数:函数的名字不能相同，但是接收的类型可以一样