package main

import (
	"card/mail"
	"card/request"
	"card/response"
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/robfig/cron"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"time"
)

var mailTemplate = "尊敬的%s,您今天的打卡结果为%s"

func main() {
	log.Println("Starting...")

	c := cron.New() // 新建一个定时任务对象
	// 每天凌晨1点
	c.AddFunc("0 0 1 * * *", func() {
		configJob()
	}) // 给对象增加定时任务
	c.Start()
	select {}

	//configJob()
}

func configJob() {
	bts, err := os.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	var conf request.Config

	err = json.Unmarshal(bts, &conf)
	if err != nil {
		panic(err)
	}

	for _, user := range conf.Users {
		job(user, conf.ReportXML)
		//fmt.Println(user, conf.ReportXML)
	}
}

func job(user request.User, reportXML string) {
	token, ok := login(user.User, user.Password)
	if !ok {
		fmt.Println("登录失败！")
	} else {
		fmt.Printf("登录成功! user: %s\n", user.User)
	}
	sessionID, ok := getRealSessionID(token)
	if !ok {
		fmt.Println("获取SessionID失败! ")
	} else {
		fmt.Printf("获取SessionID成功! SessionID:%s\n", sessionID)
	}
	ok = report(sessionID, token, reportXML)
	var err error
	if ok {
		fmt.Println("打卡成功!")
		err = mail.DoSendMail(user.Email, "成功", fmt.Sprintf(mailTemplate, user.Name, "打卡成功！如不放心可手动查看结果"))
	} else {
		fmt.Println("打卡失败!")
		err = mail.DoSendMail(user.Email, "失败", fmt.Sprintf(mailTemplate, user.Name, "打卡失败！请手动打卡，或联系老胡解决"))
	}
	if err != nil {
		fmt.Printf("邮件发送失败: %s\n", user.Email)
	}
	fmt.Printf("邮件发送成功: %s\n\n", user.Email)
}

func login(user string, password string) (string, bool) {
	agent := gorequest.New()
	var resp response.Login
	_, _, errs := agent.Post("http://stu.zstu.edu.cn/webroot/decision/login").
		Set("Connection", "keep-alive").
		Set("Content-Type", "application/json").
		Set("deviceType", "android").
		Set("Host", "stu.zstu.edu.cn").
		Set("Origin", "http://stu.zstu.edu.cn").
		Set("terminal", "H5").
		Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Mobile Safari/537.36").
		Timeout(5 * time.Second).
		Send(request.NewLoginReq(user, password)).EndStruct(&resp)

	if errs != nil {
		fmt.Println(errs)
		return "", false
	}

	return resp.Data.AccessToken, true
}

// 获取真实的sessionID
func getRealSessionID(token string) (string, bool) {
	agent := gorequest.New()
	_, text, errs := agent.Get("http://stu.zstu.edu.cn/webroot/decision/view/report?__hyperlink_depth__=1&viewlet=%252Fyewubanli%252F%25E5%2581%25A5%25E5%25BA%25B7%25E7%2594%25B3%25E6%258A%25A5.cpt&__pi__=true&op=h5_write&_replaceview_=true").
		Set("Cookie", fmt.Sprintf("fine_auth_token=%s; captchaToken=", token)).
		Set("Host", "stu.zstu.edu.cn").
		Set("Referer", "http://stu.zstu.edu.cn/webroot/decision/url/mobile?null").
		Timeout(5 * time.Second).End()
	if errs != nil {
		fmt.Println("获取sessionID失败")
		return "", false
	}

	reSessionID, _ := regexp.Compile("get sessionID[(][)] {return '(.+?)'},")
	a := reSessionID.FindStringSubmatch(text)
	return a[1], true
}

func report(sessionID string, token string, reportXML string) bool {
	agent := gorequest.New()

	resp, _, errs := agent.Post("http://stu.zstu.edu.cn/webroot/decision/view/report").
		Set("Authorization", fmt.Sprintf("Bearer %s", token)).
		Set("Cache-Control", "no-cache").
		Set("Connection", "keep-alive").
		Set("Content-Type", "application/x-www-form-urlencoded").
		Set("Cookie", fmt.Sprintf("captchaToken=; fine_auth_token=%s", token)).
		Set("deviceType", "android").
		Set("Host", "stu.zstu.edu.cn").
		Set("Origin", "http://stu.zstu.edu.cn").
		Set("sessionID", sessionID).
		Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Mobile Safari/537.36").
		Timeout(5 * time.Second).
		Send(request.NewReportAll(reportXML)).End()

	if errs != nil {
		fmt.Println(errs)
		return false
	}

	var result []response.Result
	bts, _ := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(bts, &result); err != nil {
		return false
	} else if len(result) == 0 {
		return false
	} else if !result[0].FrSubmitinfo.Success {
		return false
	}

	return true
}
