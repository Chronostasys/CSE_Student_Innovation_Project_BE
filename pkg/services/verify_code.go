package services

func IsVerifyCodeMatchToRegisterAccount(verifyCode string, email string) (IsMatch bool) {
	//re := Redis_client.Get(consts.REDIS_VERIFY_CODE_SUFFIX + email)
	//fmt.Println("\n\n", re.Val())
	//if re.Val() == verifyCode && re.Val() != "" {
	//	IsMatch = true
	//} else {
	//	IsMatch = false
	//}
	//return
	re:=Verify_codes[email]
	if len(re)!=0&&re==verifyCode{
		return true
	}else {
		return false
	}
}
func RemoveVerifyFromRedis(email string) {
	delete(Verify_codes, email)
	//Redis_client.Del(consts.REDIS_VERIFY_CODE_SUFFIX + email)
}
func StoreEmailAndVerifyCodeInRedis(verifyCode string, email string) {
	Verify_codes[email]=verifyCode
	//Redis_client.Set(consts.REDIS_VERIFY_CODE_SUFFIX+email, verifyCode, time.Hour)
}