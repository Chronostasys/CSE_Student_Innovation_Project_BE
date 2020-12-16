package models

import "github.com/jinzhu/gorm"

type AuthUser struct {
	Email                string `gorm:"primary_key"`
	Username             string
	Password             string
	Is_email_activated   bool
	Register_timestamp   int64
	Last_login_timestamp int64
	Register_ip          string
	Last_login_ip        string
	Role                 string
}
type User struct {
	gorm.Model
	sso_user_id         string
	activated_timestamp int64
	lsat_post_timestamp int64
}