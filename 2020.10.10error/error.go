package main
	/*
	有结构体Person，包含name、age、address等属性。
	定义一个函数接收结构体类型，判断传入的person的年龄是否已经成年，
	如果已成年，返回true，
	如果未成年，返回false；
	如果年龄是0或者负数，使用error返回年龄不合法等提示信息。
	 */
import (
	"errors"
	"fmt"
)

type Perspn struct {
	name string
	age int
	address string
}

func main()  {
	per :=Perspn{"LiuJiannan",-4,"江西省赣州市南康区"}
	age,err :=judgeage(per)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(age)
	}
}

func judgeage(p Perspn) (bool,error) {
	age := p.age
	if age>=18 {
		return true,nil
	}else if age>=0&&age<18 {
		return false,nil
	}else {
		err :=errors.New("年龄不合法")
		return false,err
	}
}