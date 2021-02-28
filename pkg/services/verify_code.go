package services

import (
	"fmt"
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/pkg/consts"
	"time"
)

func IsVerifyCodeMatchToRegisterAccount(verifyCode string, email string) (IsMatch bool) {
	re := Redis_client.Get(consts.REDIS_VERIFY_CODE_SUFFIX + email)
	fmt.Println("\n\n", re.Val())
	if re.Val() == verifyCode && re.Val() != "" {
		IsMatch = true
	} else {
		IsMatch = false
	}
	return
}
func RemoveVerifyFromRedis(email string) {
	Redis_client.Del(consts.REDIS_VERIFY_CODE_SUFFIX + email)
}
func StoreEmailAndVerifyCodeInRedis(verifyCode string, email string) {
	Redis_client.Set(consts.REDIS_VERIFY_CODE_SUFFIX+email, verifyCode, time.Hour)
}