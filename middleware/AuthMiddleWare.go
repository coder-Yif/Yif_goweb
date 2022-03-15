package middleware

import (
	"awesomeProject3/common"
	"awesomeProject3/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取Authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			c.JSON(http.StatusUpgradeRequired, gin.H{"code": 400, "msg": "权限不足"})
			c.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUpgradeRequired, gin.H{"code": 400, "msg": "权限不足"})
			c.Abort()
			return
		}
		userId := claims.UserId
		DB := common.InitDB()
		var user model.User
		DB.First(&user, userId)
		if user.ID == 0 {
			c.JSON(http.StatusUpgradeRequired, gin.H{"code": 400, "msg": "权限不足"})
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
