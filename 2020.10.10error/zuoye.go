package main

import (
	"errors"
	"fmt"
)

/*
	定义长度为6的数组并进行赋初始值。
	使用for循环访问0到10的数组下标上的元素。
	处理程序可能遇到的异常，处理方式主动终止程序执行，并给出异常的原因。
	 */
func main()  {
	var arr = [6]int{1,2,3,4,5,6}
	for index:= 0;index<10 ; index++ {
		if index>=len(arr) {
			fmt.Println(errors.New("打印长度已达数组长度，超过了，无法访问"))
			return
		}
		fmt.Println(arr[index])
	}
}