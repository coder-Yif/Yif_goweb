package controller

import (
	"awesomeProject3/common"
	"awesomeProject3/model"
	"awesomeProject3/utill"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	db := common.InitDB()
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	password := c.PostForm("password")
	if len(phone) != 11 {
		c.JSON(http.StatusUpgradeRequired, gin.H{"code": 422, "msg": "手机号不对"})
		return
	}
	if len(password) < 6 {
		c.JSON(http.StatusUpgradeRequired, gin.H{"code": 422, "msg": "密码太好猜了"})
		return
	}
	if len(name) == 0 {
		name = utill.RandomString(10)
	}
	if IsPhoneExist(db, phone) {
		c.JSON(http.StatusUpgradeRequired, gin.H{"code": 422, "msg": "用户已经存在"})
		return
	}
	newuser := model.User{
		Name:     name,
		Phone:    phone,
		Password: password,
	}
	db.Create(&newuser)

	log.Println(name, password, phone)
	c.JSON(200, gin.H{
		"message": "注册成功",
	})
}

func IsPhoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("Phone=?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
