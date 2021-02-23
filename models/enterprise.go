package models

import "time"

type Enterprise struct{
	ID                uint `gorm:"primary_key"`
	CreatTimestamp    time.Time
	Phone             string
	Email             string
	Name              string
	BriefIntroduction string
	Content           string
}