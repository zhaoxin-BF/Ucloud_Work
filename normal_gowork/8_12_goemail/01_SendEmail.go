/**
 * @Author: boreas.zhao email: boreas.zhao@ucloud.cn
 * @Date: 2020/8/12 5:51 下午
 * @Description:
 */

package main

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	"strconv"
	"time"
)

func SendMail(mailTo []string, subject string, body string) error {
	mailConn := map[string]string{
		"user": "1835783944@qq.com",
		"pass": "codxfwnrjuqfbhih",
		//大家一定要注意，这里的密码不是你的邮箱的密码
		//是你的邮箱申请给你的一个 --授权码---
		"host": "smtp.qq.com",
		"port": "25",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(mailConn["user"], "upupupup”")) //设置邮件发送人别名容易发送入垃圾箱
	// 这种方式可以添加别名，即“go的慢慢学习路”
	// 说明：如果是用网易邮箱账号发送，以下方法别名可以是中文，如果是qq企业邮箱，以下方法用中文别名，会报错，需要用上面此方法转码
	//m.SetHeader("From", "FB Sample"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“FB Sample”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	//m.SetHeader("From", mailConn["user"])
	m.SetHeader("To", mailTo...)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err

}

func main() {
	//定义收件人
	mailTo := []string{
		"1361412648@qq.com", //设置多个邮箱可以群发
		//"977388153@qq.com",
	}
	for i := 0; i<1000; i++{
		time.Sleep(time.Second * 1)
		go func( i int) {
			subject := fmt.Sprintln("呼叫李熙根李skr第 "+strconv.Itoa(i)+"封邮件")
			body := fmt.Sprintln("<h1>李熙根李skr第"+strconv.Itoa(i)+"次呼叫</h1>")
			err := SendMail(mailTo, subject, body)
			if err != nil {
				log.Println(err)
				fmt.Println("send fail")
				return
			}
			fmt.Println("send successfully---", i)
		}(i)
	}
	//subject := "最后一次呼叫"
	//body := "<h1>臭男人 over 了</h1>"
	//err := SendMail(mailTo, subject, body)
	//if err != nil {
	//	log.Println(err)
	//	fmt.Println("send fail")
	//	return
	//}

	//fmt.Println("send successfully")

}