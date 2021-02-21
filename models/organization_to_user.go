package models

import "time"

type OrganizationToUser struct{
	ID                 uint `gorm:"primary_key"`
	OrganizationID     uint
	UserID             uint
	JoinAtTimestamp    time.Time//用户加入组织的时间
}
//这个表主要用于建立组织和用户间的多对多关系