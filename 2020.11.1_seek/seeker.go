package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	/*
		断点续传:
			文件传递:文件复制

			"b.jpg"

			复制到当前的工程下
		思路:边复制,边记录复制的总量
	*/
	srcFile := "b.jpg"
	destFile := srcFile[strings.LastIndex(srcFile, "/")+1:]
	fmt.Println(destFile)
	tempFile := destFile + "temp.txt"
	fmt.Println(tempFile)

	file1, err := os.Open(srcFile)
	HandLeErr(err)
	file2,err := os.OpenFile(destFile,os.O_CREATE|os.O_WRONLY,os.ModePerm)
	HandLeErr(err)
	file3,err := os.OpenFile(tempFile,os.O_CREATE|os.O_RDWR,os.ModePerm)
	HandLeErr(err)

	defer file1.Close()
	defer file2.Close()

	//step1:先读取临时文件中的数据,再seek
	file3.Seek(0,io.SeekStart)
	bs := make([]byte,100,100)
	n1,err := file3.Read(bs)
	HandLeErr(err)
	countStr := string(bs[:n1])
	count,err := strconv.ParseInt(countStr,10,64)
	//HandLeErr(err)
	fmt.Println(count)

	//step2:设置读,写的位置:
	file1.Seek(count,io.SeekStart)
	file2.Seek(count,io.SeekStart)
	data := make([]byte,1024,1024)
	n2 := -1 //读取的数据量
	n3 := -1 //写出的数据量
	total := int(count) //读取的总量

	//step3 := 复制文件
	for{
		n2,err = file1.Read(data)
		if err == io.EOF || n2 == 0 {
			fmt.Println("文件复制完毕。。")
			file3.Close()
			os.Remove(tempFile)
			break
		}
		n3,err = file2.Write(data[:n2])
		total += n3
	}

	//将复制的总量,存储到临时文件中
	file3.Seek(0,io.SeekStart)
	file3.WriteString(strconv.Itoa(total))

	fmt.Printf("total:%d\n", total)

	//假装断电
	//if total > 8000{
	//	panic("假装断电了。。。")
	//}

}
func HandLeErr(err error)  {
	if err != nil {
		fmt.Println(err)
		return
	}
}