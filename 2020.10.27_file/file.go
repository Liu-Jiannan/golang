package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main()  {
	/*
	文件操作:
	1.路径:
		相对路径:relative
			ab.txt
			相当于当前工程
		绝对路径:absolute
			C:\\User\\GoProjects\\src\\2020.10.27_file\\aa.txt
		.表示当前目录
		..上一层目录
	2.创建文件夹,如果文件夹存在，创建失败
		os.Mkdir(),创建一层
		os.MkdirAll(),创建多层
	3.创建文件:Create采用模式0666(任何人都可读写,不可执行) 创建一个名为name的文件，如果文件已存在会截断它(为空文件)
		 os.Create() 创建文件
	4.打开文件:让当前的程序，和指定的文件之间建立一个连接
		os.Open(filename)
		os.openFile(filename,mode,perm)
	5.关闭文件.
		file4.Close()
	6.删除文件或文件夹:慎用，慎用，再慎用
		os.Remove(""),删除文件和空目录
		os.RemoveAll(),删除所有

	*/
	//1.路径
	fileName1:="C:\\User\\GoProjects\\src\\2020.10.27_file\\aa.txt"
	fileName2:="aa.txt"
	fmt.Println(filepath.IsAbs(fileName1))
	fmt.Println(filepath.IsAbs(fileName2))
	//获取绝对路径
	fmt.Println(filepath.Abs(fileName1))
	fmt.Println(filepath.Abs(fileName2))

	//获取父目录
	fmt.Println("获取父目录:",path.Join(fileName1,".."))

	//创建目录
	//err:=os.Mkdir(".\\b",os.ModePerm)
	//if err!=nil {
	//	fmt.Println("err",err)
	//	return
	//}
	//fmt.Println("文件夹创建成功")
	//err=os.MkdirAll(".\\c\\d",os.ModePerm)
	//if err!=nil {
	//	fmt.Println("err:",err)
	//	return
	//}
	//fmt.Println("多层目录创建成功")

	//3.创建文件:Create采用模式0666(任何人都可读写,不可执行) 创建一个名为name的文件，如果文件已存在会截断它(为空文件)
	//file1,err:=os.Create(".\\b\\bb.txt")
	//if err!=nil {
	//	fmt.Println("err:",err)
	//	return
	//}
	//fmt.Println(file1)

	//4.打开文件
	//file3,err:=os.Open(fileName1)//只读的
	//if err!=nil{
	//	fmt.Println("err:",err)
	//	return
	//}
	//fmt.Println(file3)
	//file4,err:=os.OpenFile(fileName1,os.O_RDONLY|os.O_WRONLY,os.ModePerm)
	/*
	第一个参数: 文件名称
	第二个参数：文件的打开方式
	第三个参数: 文件的权限:文件不存在创建文件，需要指定权限
	 */
	//if err!=nil {
	//	fmt.Println("err:",err)
	//	return
	//}
	//fmt.Println(file4)
	//5.关闭文件
	//file4.Close()

	//6.删除文件或文件夹:
	//err := os.Remove("C:\\User\\GoProjects\\src\\2020.10.27_file\\b")
	//if err != nil {
	//	fmt.Println("err:",err)
	//	return
	//}
	//fmt.Println("文件删除成功!!!")
}
