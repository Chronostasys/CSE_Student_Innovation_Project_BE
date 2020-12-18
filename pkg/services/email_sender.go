package services

import (
	"gopkg.in/gomail.v2"
)

func InviteMemberToTeam(email_from string,team_id int,description string,email_to string){
	m := gomail.NewMessage()
	// 发邮件的地址
	m.SetHeader("From", "shenboyu2020@163.com")
	// 给谁发送，支持多个账号
	m.SetHeader("To", email_to)
	// 抄送谁
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	// 邮件标题
	m.SetHeader("Subject", "您正被邀请加入团队")
	// 邮件正文，支持 html
	//这里等着前端给我写html啦
	Owner_email:=email_from
	email_accept:=email_to
	url:=CreatEmailContentUrl(Owner_email,team_id,description,email_accept)
	m.SetBody("text/html", url+description)
	// 附件
	//m.Attach("/home/Alex/lolcat.jpg")
	// stmp服务，端口号，发送邮件账号，发送账号密码
	d := gomail.NewDialer("smtp.163.com", 25, "shenboyu2020@163.com", "WOEVMSPVCYASVDNSk")
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
