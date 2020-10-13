package main

import "./utils"

func main()  {
	utils.Count()
	/*
	init()	main():
	相同点:
	这两个函数都是go语言中的保留函数，init()用于初始化信息。main用于作为程序的入口
	init()与main()不能有任何的参数与返回值
	该函数只能由go程序自动调用,不可以被引用
	不同点:
	init可以应用与任意包中,且只能定义一个
	main函数只能用于main包中,且只能定义一个.

	两个函数的执行顺序:
	在main包中的go文件默认总是会被执行。
	对同一个go文件的init()调用顺序是从上到下的
	在同一个包中的不同文件,将按照文件名字符串进行"从小到大"排序调用
	对于不同的包，如果不相互依赖的话，按照main包中import的顺序调用其包中的函数.
				如果相互依赖，调用顺序为最后被依赖的最先被初始化
	 */

	/*
	init()函数和main()函数
	1.

	 */
}