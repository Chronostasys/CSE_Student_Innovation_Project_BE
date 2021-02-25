package models

import "time"

type Enterprise struct{
	ID                uint `gorm:"primary_key"`
	CreatTimestamp    time.Time
	Address           string
	Phone             string
	Email             string
	Name              string
	BriefIntroduction string
	Content           string
	CreaterEmail	  string
	Image			  string
}