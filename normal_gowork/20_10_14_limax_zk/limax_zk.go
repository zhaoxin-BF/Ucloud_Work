/**
 * @Author: boreas.zhao email: boreas.zhao@ucloud.cn
 * @Date: 2020/10/14 5:28 下午
 * @Description: limax zookeeper 变更脚本
 */

package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main(){

	//路径
	path := "/Users/boreas/Workspace/verify_set/HNSSN08"                     //测试
	//path := "/Users/boreas/Workspace/limax/udisk/region"

	//1、给每个set 创建文件并写入数据default
	r := regexp.MustCompile("[0-9]{4}")

	set, _ := ioutil.ReadDir(path)
	for _, val := range set{
		if r.MatchString(val.Name()) {
			fmt.Println("集群！")
		}
	}





	//2、读取每个region zookeeper 信息

	//3、kv json 形式写入 zookeeper
}