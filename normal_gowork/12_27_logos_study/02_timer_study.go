package main

import (
	"fmt"
	"time"
)

func main() {
	//1、时间转换为字符串
	timeEnd := time.Now().Format("2006年01月02日 15时04分05秒")
	fmt.Println(timeEnd) //输出：2019-12-27 10:59:36 or 2019年12月27日 11时06分33秒
	fmt.Printf("%T\n", timeEnd)

	//输出
	// 2019年12月27日 11时09分29秒
	// string

	//2、字符串转换为时间格式
	layout := "2006-01-02 15:04:05"
	str := "2016-07-25 11:45:26"
	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)
	fmt.Printf("%T\n", t)

	//输出
	// 2016-07-25 11:45:26 +0000 UTC
	// time.Time

}
//普通unix时间转换
func unix(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04:05")
}

// nano 纳秒转换
func nano (timestamp float64, nsec int64) string {
	//纳秒没什么用 前段不显示 直接将小数舍弃转化为int64
	tm := time.Unix(int64(timestamp), nsec)
	return tm.Format("2006-01-02 15:04:05")
}
