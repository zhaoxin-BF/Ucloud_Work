/**
 * @Author: boreas.zhao email: boreas.zhao@ucloud.cn
 * @Date: 2020/7/30 8:15 下午
 * @Description:
 */

package main

import (
	"fmt"
	"io/ioutil"
	"syscall"
)

func main() {

	syscall.Umask(0000)

	//路径+文件名
	err := ioutil.WriteFile("12_4_study_day3_01_json/aaaass", []byte("ffsfsf"), 0777)
	if err != nil {
		fmt.Printf("ioutil.WriteFile failure, err=[%v]\n", err)
	}
}