package router

import (
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/models"
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/pkg/consts"
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/pkg/services"
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "helloworld",
	})
	return


}
func Signup(context *gin.Context) {
	email := strings.ToLower(context.PostForm("email"))
	rsa_password := context.PostForm("password")
	verifyStr := context.PostForm("verify_code")
	name :=context.PostForm("name")
	//password in request is encrypted by rsa,so I should decrypt it first
	password := rsa_password

	if services.IsEmailRegistered(email) {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "该用户已经注册",
		})
		return
	}
	if services.IsVerifyCodeMatchToRegisterAccount(verifyStr, email) {
		services.RemoveVerifyFromRedis(email)
		passwordhash := util.HashWithSalt(password)
		user := models.AuthUser{
			Email: email, Password: passwordhash,Username: name,
			Register_timestamp: util.GetTimeStamp(), Is_email_activated: true,
			Role: consts.USER}
		services.AddUserWithoutCheck(user)
		token, _ := util.GenerateToken(email, consts.USER)
		context.SetCookie(consts.COOKIE_NAME, token, consts.EXPIRE_TIME_TOKEN, "/", "localhost", false, true)
		context.JSON(http.StatusOK, gin.H{
			"email": user.Email,
			"msg":   "注册成功",
		})
		return
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码错误，请重新输入或者获取",
		})
		return
	}

}
func Login(context *gin.Context) {
	email := strings.ToLower(context.PostForm("email"))
	password := context.PostForm("password")

	role := services.GetRoleWhileLogin(email)
	token, _ := util.GenerateToken(email, role)

	if !services.IsEmailRegistered(email) {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "该账户未注册，请首先注册",
		})
		return
	}

	if isMacth := services.IsEmailAndPasswordMatch(email, password); isMacth {
		context.SetCookie(consts.COOKIE_NAME, token, 1000, "/", "localhost", false, true)
		context.JSON(http.StatusOK, gin.H{
			"msg": "登陆成功!",
		})
		return
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "用户名密码不匹配",
		})
		return
	}
}
func SendVerifyCode(context *gin.Context) {
	email := strings.ToLower(context.PostForm("email"))
	verifyCode := util.GenerateVerifyCode(consts.VERIFYCODE_LENGTH)
	services.StoreEmailAndVerifyCodeInRedis(verifyCode, email)

	if sendEmailSuccessful := services.SendEmail(email, verifyCode); sendEmailSuccessful {
		context.JSON(http.StatusOK, gin.H{
			"msg": "成功发送验证码，请注意查收",
		})
		return
	} else {
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg": "发送验证码失败，请重试，如果多次失败，请联系管理员",
		})
		return
	}
}
func ChangePassword_Email(context *gin.Context){
	email := strings.ToLower(context.PostForm("email"))
	if !services.IsEmailRegistered(email) {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "该账户未注册，请首先注册",
		})
		return
	}else{
		context.JSON(http.StatusOK, gin.H{
			"msg": "账户已注册",
		})
		return
	}
}
func ChangePasswordVerify_code(context *gin.Context){
	email := strings.ToLower(context.PostForm("email"))
	verifyStr := context.PostForm("verify_code")
	if services.IsVerifyCodeMatchToRegisterAccount(verifyStr, email) {
		services.RemoveVerifyFromRedis(email)
		context.JSON(http.StatusOK, gin.H{
			"msg":   "验证码正确",
		})
		return
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码错误，请重新输入或者获取",
		})
		return
	}
}
func ChangePassword(context *gin.Context) {
	email := strings.ToLower(context.PostForm("email"))
	verifyCode := context.PostForm("verify_code")
	newPassword := context.PostForm("new_password")

	if services.IsVerifyCodeMatchToRegisterAccount(verifyCode, email) {
		passwordHash := util.HashWithSalt(newPassword)
		err := services.ResetUserPassword(email, passwordHash)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		} else {
			context.JSON(http.StatusOK, gin.H{
				"msg": "修改密码成功",
			})
			return
		}
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码过期或不正确",
		})
		return
	}
}
