/**
 * @Author: boreas.zhao email: boreas.zhao@ucloud.cn
 * @Date: 2020/7/31 10:26 上午
 * @Description:
 */

package main
import (
	"fmt"
	"io/ioutil"
	"os/exec"
)
func main() {
	err, ret := ExecCmd()
	if err != nil {
		fmt.Println("出错了！")
		return
	}
	fmt.Println("dsfsdfs",ret)
}

func ExecCmd() (err error, result string){
	cmd := exec.Command("/bin/bash", "-c", `python test.py`)
	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		return err, result
	}
	//执行命令
	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err,", err)
		return err, result
	}
	//读取所有输出
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll Stdout:", err.Error())
		return err, result
	}
	if err := cmd.Wait(); err != nil {
		fmt.Println("wait:", err.Error())
		return err, result
	}
	fmt.Printf("stdout:\n\n %s", bytes)
	result = string(bytes)
	return err, result
}