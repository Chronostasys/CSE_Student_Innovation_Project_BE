package models
type AuthUser struct {
	Email                string `gorm:"primary_key"`
	Username             string
	Password             string
	RegisterTimestamp    int64
	LastLoginTimestamp   int64
	Role                 string
}

