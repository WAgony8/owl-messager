package model

// MailSenderInfo 邮件发送账号
type MailSenderInfo struct {
	Pk
	Code       string `orm:"column(code);unique"`
	Name       string `orm:"column(name)"`
	Host       string `orm:"column(host)"`
	Port       int    `orm:"column(port)"`
	Email      string `orm:"column(email)"`
	EmailAlias string `orm:"column(email_alias)"` // 20211028 新增列 发件人的子邮箱 “from EmailAlias”
	Alias      string `orm:"column(alias);null"`
	Passwd     string `orm:"column(password)"`
	Default    int    `orm:"column(in_default)"`

	TableChangeInfo
}