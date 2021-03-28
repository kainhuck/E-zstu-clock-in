package mail

import "testing"

func TestMail(t *testing.T) {
	e := DoSendMail("1628321264@qq.com", "成功", "晚上好阿")
	if e != nil {
		panic(e)
	}
}
