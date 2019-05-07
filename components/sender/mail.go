package sender

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/go-gomail/gomail"
	"strconv"
)

type Mailer struct {
	Address string
	Name    string
}

type MailContent struct {
	From    Mailer
	To      []Mailer
	Cc      []Mailer
	Subject string
	Body    string
	Attach  string
}

var (
	content *MailContent
	dialer  *gomail.Dialer
)

func SendMail(mc *MailContent) {
	registerMailDialer()
	m := gomail.NewMessage()
	if (mc.From == Mailer{}) {
		m.SetHeader("From", m.FormatAddress(
			beego.AppConfig.String("mail::username"),
			beego.AppConfig.String("mail::name")),
		)
	} else {
		m.SetHeader("From", m.FormatAddress(mc.From.Address, mc.From.Name))
	}

	if len(mc.To) == 0 {
		logs.Error("Have no receiver")
		return
	}

	var toMails []string

	for _, to := range mc.To {
		toMails = append(toMails, m.FormatAddress(to.Address, to.Name))
	}

	m.SetHeader("To", toMails...)

	if len(mc.Cc) != 0 {
		var ccMails []string
		for _, cc := range mc.Cc {
			ccMails = append(ccMails, m.FormatAddress(cc.Address, cc.Name))
		}
		m.SetHeader("Cc", ccMails...)
	}

	m.SetHeader("Subject", mc.Subject)
	m.SetBody("text/html", template(mc.Body))
	if mc.Attach != "" {
		m.Attach(mc.Attach)
	}

	if err := dialer.DialAndSend(m); err != nil {
		logs.Error("Mail send failed")
	}
}

func template(body string) (html string) {
	html = "template:" + body
	return
}

func registerMailDialer() {
	if dialer != nil {
		return
	}

	host := beego.AppConfig.String("mail::host")
	portStr := beego.AppConfig.String("mail::port")
	port, _ := strconv.Atoi(portStr)
	username := beego.AppConfig.String("mail::username")
	password := beego.AppConfig.String("mail::password")
	dialer = gomail.NewDialer(host, port, username, password)
}

func NewContent() *MailContent {
	content = &MailContent{}
	return content
}

func (mc *MailContent) SetTo(to map[string]string) *MailContent {
	var tos []Mailer
	for name, address := range to {
		tos = append(tos, Mailer{
			Address: address,
			Name:    name,
		})
	}
	content.To = tos

	return content
}

func (mc *MailContent) SetCc(cc map[string]string) *MailContent {
	var ccs []Mailer
	for name, address := range cc {
		ccs = append(ccs, Mailer{
			Address: address,
			Name:    name,
		})
	}
	content.Cc = ccs

	return content
}

func (mc *MailContent) SetFrom(address, name string) *MailContent {
	content.From.Address = address
	content.From.Name = name

	return content
}

func (mc *MailContent) SetSubject(subject string) *MailContent {
	content.Subject = subject

	return content
}

func (mc *MailContent) SetBody(body string) *MailContent {
	content.Body = body

	return content
}
