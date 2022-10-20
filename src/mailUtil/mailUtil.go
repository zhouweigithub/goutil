package mailutil

import (
	"errors"
	errutil "goutil/errUtil"
	logutil "goutil/logUtil"
	randutil "goutil/randUtil"
	"strings"
	"time"

	"gopkg.in/gomail.v2"
)

// MailHelper 邮件功能
type MailHelper struct {
	// ServerHost 邮箱服务器地址，如腾讯企业邮箱为smtp.qq.com
	ServerHost string
	// ServerPort 邮箱服务器端口，如腾讯企业邮箱为465
	ServerPort int
	// FromEmail　发件人邮箱地址
	FromEmail string
	// FromPasswd 发件人邮箱密码（注意，这里是明文形式），TODO：如果设置成密文？
	FromPasswd string

	msg    *gomail.Message
	dialer *gomail.Dialer
}

// Init 初始化邮件选项
func (m *MailHelper) Init(host string, port int, from string, pwd string) {
	defer errutil.CatchError()
	m.ServerHost = host
	m.ServerPort = port
	m.FromEmail = from
	m.FromPasswd = pwd

	m.msg = gomail.NewMessage()
	m.dialer = gomail.NewDialer(m.ServerHost, m.ServerPort, m.FromEmail, m.FromPasswd)

	// 发件人
	// 第三个参数为发件人别名，如"李大锤"，可以为空（此时则为邮箱名称）
	m.msg.SetAddressHeader("From", from, "客服助理")
}

// SendEmail 同步发送邮件（需要先调用Init()方法）
func (m *MailHelper) SendEmail(subject, body, tors string) error {
	defer errutil.CatchError()
	// 主题
	m.msg.SetHeader("Subject", subject)
	// 正文
	m.msg.SetBody("text/html", body)

	var toArray []string
	if len(tors) == 0 {
		var msg = "邮件未发送，未配置收件人"
		logutil.Error(msg)
		return errors.New(msg)
	}
	for _, tmp := range strings.Split(tors, ",") {
		toArray = append(toArray, strings.TrimSpace(tmp))
	}
	// 收件人可以有多个，故用此方式
	m.msg.SetHeader("To", toArray...)

	// 抄送
	// m.msg.SetHeader("Cc", toArray...)

	// 发送
	err := m.dialer.DialAndSend(m.msg)
	if err != nil {
		//如果出错，则重试n次
		for i := 0; i < 2; i++ {
			//等待时间
			var waitSeconds = randutil.GetRandInt(5, 10)
			time.Sleep(time.Second * time.Duration(waitSeconds))
			err = m.SendEmail(subject, body, tors)
			if err == nil {
				break
			}
		}
		if err != nil {
			logutil.Error("邮件发送失败：" + err.Error())
		}
	}

	return err
}

// SendEmailAsync 异步发送邮件（需要先调用Init()方法）
func (m *MailHelper) SendEmailAsync(subject, body, tors string) {
	go m.SendEmail(subject, body, tors)
}
