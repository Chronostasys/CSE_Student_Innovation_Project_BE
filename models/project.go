package models

import "time"

type Project struct{
	ID                uint `gorm:"primary_key"`
	CreatTimestamp    time.Time
	Phone             string
	Email             string
	Name              string
	BriefIntroduction string
	Content           string
}