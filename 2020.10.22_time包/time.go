package main

import (
	"fmt"
	"time"
)

/**
	time包:
		1年=365天	day
		1天=24小时	hour
		1小时=60分钟	minute
		1分钟=60秒	secend
		1秒种=1000毫秒	millisecond
		1毫秒=1000微秒	microsecond->us
		1微秒=1000纳秒	niaosecond->ns
		1纳秒=1000皮秒	picosecond->ps
 */
func main()  {
	//1.获取当前的时间
	t1 := time.Now()
	fmt.Printf("%T\n",t1)
	fmt.Println(t1)

	//2。设置指定的时间
	t2 := time.Date(2000,7,15,10,40,2,2,time.Local)
	fmt.Println(t2)

	//3.time->string之间的转换
	/*
	t1.Format("格式模板")-->string
		模板的日期必须是固定:06-1-2-3-4-5
	 */
	s1 := t1.Format("2006年1月2日 15:04:05")
	fmt.Println(s1)

	s2 := t1.Format("2006/01/02")
	fmt.Println(s2)

	//4.string->time
	/*
	time.Parse("模板",str)-->time,err
	 */
	s3 := "1999年10月10日"//string
	t3,err := time.Parse("2006年01月02日",s3)
	if err!= nil {
		fmt.Println("err:",err)
	}
	fmt.Println(t3)
	fmt.Printf("%T\n",t3)

	//4.根据当前时间，获取指定的内容
	year,month,day:=t1.Date()
	fmt.Println(year,month,day)

	hour,min,sec := t1.Clock()
	fmt.Println(hour,min,sec)

	fmt.Println(t1.Year())
	fmt.Println(t1.Month())
	fmt.Println(t1.Day())
	fmt.Println(t1.Hour())
	fmt.Println(t1.Minute())
	fmt.Println(t1.Second())
	fmt.Println(t1.Nanosecond())

	fmt.Println(t1.Weekday())	//星期几


	}