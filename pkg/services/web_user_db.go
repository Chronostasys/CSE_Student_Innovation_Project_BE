package services
import "github.com/Pivot-Studio/CSE_Student_Innovation_Project/models"

func GetUserNameByEmail(ownerEmail string)(userName string){
	var authUser models.AuthUser
	db.Where("Email = ?",ownerEmail).First(&authUser)

	userName = authUser.Username
	return
}