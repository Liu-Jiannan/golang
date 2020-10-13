package utils

import "fmt"

//首字母大写，代表可以被导出,表示为公共的,这个包外的其它文件可以通过导包语句可导出使用
func Count(){
	fmt.Println("utils包下的Count()函数")
}