package main

import (
	"fmt"
	"log"
	"io/ioutil"
)

func main()  {
	/*
	遍历文件夹:
	 */
	dirname := "/User/GoProjects/src"
	listFiles(dirname,0)
}

func listFiles(dirname string,level int)  {
	//level用来记录当前递归的层次，生成带有层次感的空格
	s := "|--"
	for i:=0; i<level; i++ {
		s = "|"+s
	}
	fileInfos,err :=ioutil.ReadDir(dirname)
	if err!=nil {
		log.Fatal()
	}
	for _,fi:= range fileInfos{
		filename:=dirname+"/"+fi.Name()
		fmt.Printf("%s%s\n",s,filename)
		if fi.IsDir() {
			//递归调用方法
			listFiles(filename,level+1)
		}
	}
}