package main

import (
	"fmt"
	"os"
)

func main()  {
	/*
	FileInfo:文件信息
		inferface
			Name(),文件名
			Size(),文件大小(字节)
			IsDir(),是否目录
			ModTime(),修改时间
			Mode(),权限

	 */
	FileInfo,err := os.Stat("C:\\User\\GoProjects\\src\\2020.10.27_file\\aa.txt")
	if err != nil {
		fmt.Println("err:",err)
		return
	}
	fmt.Printf("%T\n",FileInfo)
	//文件名
	fmt.Println(FileInfo.Name())
	//文件大小
	fmt.Println(FileInfo.Size())
	//是否是目录(文件夹)
	fmt.Println(FileInfo.IsDir())
	//修改时间
	fmt.Println(FileInfo.ModTime())
	//权限
	fmt.Println(FileInfo.Mode())//-rw-rw-rw-
}
