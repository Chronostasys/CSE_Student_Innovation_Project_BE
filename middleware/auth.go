package middleware

import (
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/pkg/consts"
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/pkg/services"
	"github.com/Pivot-Studio/CSE_Student_Innovation_Project/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString, err := context.Request.Cookie(consts.COOKIE_NAME)
		if err != nil {
			context.Abort()
			context.JSON(http.StatusUnauthorized, gin.H{
				"msg": "未认证，请先登录",
			})
			return
		}
		claim, err := util.GetClaimFromToken(tokenString.Value)
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