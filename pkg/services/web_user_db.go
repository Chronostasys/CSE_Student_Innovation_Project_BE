package services
import "github.com/Pivot-Studio/Authorization-Template/ATmodels"

func GetUserNameByEmail(ownerEmail string)(userName string){
	var authUser ATmodels.AuthUser
	db.Where("Email = ?",ownerEmail).First(&authUser)

	userName = authUser.Username
	return
}