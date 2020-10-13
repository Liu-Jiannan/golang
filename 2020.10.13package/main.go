package main

//			包可以取别名进行导入
/*
	import(
	p1 "package1"
	p2 "package2"
)
	使用时:别名操作，调用包函数时前缀变成了我们的别名
*/

import (
	p "./person"
	"./utils"
	"fmt"
) //相对路径

import "./utils/timeutils"

import "./pk1"

func main()  {
	/*
	关于包的使用:
	1.一个目录下的统计文件归属一个包，package的声明要一致
	2.package声明的包和对应的目录名可以不一致，但习惯上还是写成一致的
	3.包可以嵌套
	4.同包下的函数不需要导入包，可以直接使用
	5.main包，main()函数所在的包，其它的包不能使用
	6.导入包的时候，路径要从package开始写


	 */
	utils.Count()
	//包嵌套
	timeutils.PrintTime()
	pk1.MyTest1()
	utils.MyTest2()
	fmt.Println("--------------")

	p1 := p.Person{"李小花",30,"女"}
	fmt.Println(p1.Name,p1.Age,p1.Sex)
}


