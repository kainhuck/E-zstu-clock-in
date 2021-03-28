package mail

import (
	"gopkg.in/gomail.v2"
	"strconv"
)


func SendMail(mailTo []string, subject string, body string) error {

	mailConn := map[string]string{
		"user": "123456789@qq.com", // 邮件发送者的地址
		"pass": "dcrdyvbuidsxcvn",  // qq邮箱填授权码，百度一下获取方式。
		"host": "smtp.qq.com",       // 发送将邮件发送给腾讯的smtp邮件服务器
		"port": "465",               // 发送邮件使用的端口
	}
	port, _ := strconv.Atoi(mailConn["port"])

	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(mailConn["user"], "E浙理自动打卡结果通知"))
	m.SetHeader("To", mailTo...)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
	err := d.DialAndSend(m)
	return err
}

func DoSendMail(stuEmail, subject, body string) (e error) {
	mailTo := []string{stuEmail}
	err := SendMail(mailTo, subject, body)

	if err != nil {
		e = err
		return e
	}
	return nil
}
