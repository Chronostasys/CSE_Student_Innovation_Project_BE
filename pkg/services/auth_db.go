package services

import (
	"errors"
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/models"
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/pkg/util"
	"golang.org/x/crypto/bcrypt"

)

//Email is the chiper and pwd is plaintext which will be hashed in function
func IsEmailAndPasswordMatch(email string, pwd string) (isEmailAndPasswordMatch bool) {
	var user models.AuthUser
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

//encrypted_email is the cipher after encrypting email,pay attention to
//passing email with encryption
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

//encrypted_email is the cipher after encrypting email,pay attention to
//passing email with encryption
func VerifyEmailActive(encrypted_email string) {
	var user models.AuthUser
	db.Where("email=?", encrypted_email).Find(&user)
	user.Is_email_activated = true
	db.Save(&user)
	return
}
func JoinHoleTime(email_aes string) (day int) {
	var auth models.AuthUser
	db.Where("email=?", email_aes).Find(&auth)
	sec := util.GetTimeStamp() - auth.Register_timestamp
	day = int(sec / (60 * 60 * 24))
	return day
}

func ChangePassword(email_aes string, oldPassword string, newPassword string) (result bool) {
	var user models.AuthUser
	db.Where("email=?", email_aes).Find(&user)
	if user.Password != util.HashWithSalt(oldPassword) {
		return false
	} else {
		user.Password = util.HashWithSalt(newPassword)
		db.Save(&user)
		return true
	}
}
