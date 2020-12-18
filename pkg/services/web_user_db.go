package services
import "github.com/Pivot-Studio/Authorization-Template/models"

func GetUserNameByEmail(ownerEmail string)(userName string){
	var authUser models.AuthUser
	db.Where("Email = ?",ownerEmail).First(&authUser)

	userName = authUser.Username
	return
}