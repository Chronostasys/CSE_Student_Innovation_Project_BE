package services

import (
	_ "database/sql"
	"fmt"
	atmodels "github.com/Pivot-Studio/Authorization-Template/models"
	"github.com/Pivot-Studio/Authorization-Template/pkg/util"
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func dsn(settings atmodels.DbSettings) string {
	// https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
	// Add ?parseTime=true
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8mb4,utf8", settings.Username, settings.Password, settings.Hostname, settings.Dbname)
}
func init(){
	databaseInit()
}
func databaseInit() {
	conf := util.ReadSettingsFromFile("Config.json")
	settings := conf.DbSettings
	connStr := dsn(settings)

	var err1 error
	db, err1 = gorm.Open("mysql", connStr)
	util.CheckError(err1)

	var blog models.Blog
	var team models.Team
	var user atmodels.AuthUser
	var invite models.Invite
	var temp []interface{}

	temp = append(temp,&blog,&team,&invite,&user)
	util.CreateTableIfNotExist(db, temp)
}