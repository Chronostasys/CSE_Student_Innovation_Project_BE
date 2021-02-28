package router

import (
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/models"
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/pkg/consts"
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/pkg/services"
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func HelloWorld(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "HelloWorld",
	})
	return
}
func Signup(context *gin.Context) {
	email := strings.ToLower(context.PostForm("email"))
	rsaPassword := context.PostForm("password")
	verifyStr := context.PostForm("verify_code")
	name :=context.PostForm("name")
	//password in request is encrypted by rsa,so I should decrypt it first
	password := rsaPassword

	if services.IsEmailRegistered(email) {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "该用户已经注册",
		})
		return
	}
	if services.IsVerifyCodeMatchToRegisterAccount(verifyStr, email) {
		services.RemoveVerifyFromRedis(email)
		passwordHash := util.HashWithSalt(password)
		user := models.AuthUser{
			Email: email, Password: passwordHash,Username: name,
			RegisterTimestamp: util.GetTimeStamp(),
			Role: consts.USER}
		services.CreateUser(user)
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

	if isMatch := services.IsEmailAndPasswordMatch(email, password); isMatch {
		context.JSON(http.StatusOK, gin.H{
			"token":token,
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
func ChangePasswordByEmail(context *gin.Context){
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
func ChangePasswordVerifyCode(context *gin.Context){
	email := strings.ToLower(context.PostForm("email"))
	verifyStr := context.PostForm("verify_code")
	if services.IsVerifyCodeMatchToRegisterAccount(verifyStr, email) {
		//services.RemoveVerifyFromRedis(email)
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
		services.RemoveVerifyFromRedis(email)
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
func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString:= context.GetHeader("token")
		claim, err := util.GetClaimFromToken(tokenString)
		if err != nil {
			context.Abort()
			context.JSON(http.StatusUnauthorized, gin.H{
				"msg": "未认证，请先登录",
			})
			return
		}
		email := claim.(jwt.MapClaims)["email"].(string)

		//using assert is very dangerous
		tokenTimeStamp := claim.(jwt.MapClaims)["timeStamp"].(float64)
		time := util.GetTimeStamp() - int64(tokenTimeStamp)
		if time > consts.EXPIRE_TIME_TOKEN {
			context.Abort()
			context.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Token过期，请重新登录",
			})
		}

		if services.IsEmailRegistered(email) {
			context.Next()
		} else {
			context.Abort()
			context.JSON(http.StatusUnauthorized, gin.H{
				"msg": "无效token，请重新登录",
			})
			return
		}
	}
}
func GetMyselfInfo(context *gin.Context){
	email,_:=util.GetEmailFromToken(context)
	user,_:=services.GetUserByEmail(email)
	context.JSON(http.StatusOK,gin.H{
		"name":user.Username,
		"email":user.Email,
	})
	return
}