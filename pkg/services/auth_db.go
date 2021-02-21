package services

import (
	"errors"
	"github.com/Pivot-Studio/Authorization-Template/ATmodels"
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/models"
	"golang.org/x/crypto/bcrypt"
)

func IsEmailRegistered(encrypted_email string) (IsRegistered bool) {
	var user models.AuthUser
	db.Where("email = ?", encrypted_email).Find(&user)
	if (user == models.AuthUser{}) {
		IsRegistered = false
	} else {
		IsRegistered = true
	}
	return IsRegistered
}
func CreateUser(holeuser models.AuthUser) {
	db.Create(&holeuser)
}
func GetRoleWhileLogin(email_aes string) (role string) {
	var user ATmodels.AuthUser
	db.Where("email=?", email_aes).Find(&user)
	role = user.Role
	return role
}
func IsEmailAndPasswordMatch(email string, pwd string) (isEmailAndPasswordMatch bool) {
	var user ATmodels.AuthUser
	db.Where("email = ?", email).Find(&user)
	result := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pwd))
	if result == nil {
		isEmailAndPasswordMatch = true
	} else {
		isEmailAndPasswordMatch = false
	}
	return
}
func ResetUserPassword(email string, passwordHash string) (err error) {
	var doc models.AuthUser
	db.Where("email = ?", email).Find(&doc)
	if doc.Email == "" {
		return errors.New("No such user in database")
	}
	doc.Password = passwordHash
	db.Save(&doc)
	return nil
}
func GetUserByEmail(email string)(user models.AuthUser,err error){
	err=db.Where("email=?",email).Find(&user).Error
	return
}