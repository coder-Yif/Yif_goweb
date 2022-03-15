package controller

import (
	"awesomeProject3/common"
	"awesomeProject3/dto"
	"awesomeProject3/model"
	"awesomeProject3/response"
	"awesomeProject3/utill"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	db := common.InitDB()
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	password := c.PostForm("password")
	if len(phone) != 11 {
		response.Response(c, http.StatusUpgradeRequired, 422, nil, "手机号不对")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUpgradeRequired, 422, nil, "密码太好猜了")
		return
	}
	if len(name) == 0 {
		name = utill.RandomString(10)
	}
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusUpgradeRequired, 422, nil, "加密失败")
		return
	}
	if IsPhoneExist(db, phone) {
		response.Response(c, http.StatusUpgradeRequired, 422, nil, "用户已经存在")
		return
	}
	newuser := model.User{
		Name:     name,
		Phone:    phone,
		Password: string(hasedPassword),
	}
	db.Create(&newuser)
	log.Println(name, password, phone)
	response.Success(c, http.StatusUpgradeRequired, nil, "注册成功")
}
func Login(c *gin.Context) {
	db := common.InitDB()
	phone := c.PostForm("phone")
	password := c.PostForm("password")
	var user model.User
	db.Where("Phone=?", phone).First(&user)
	if user.ID == 0 {

		response.Response(c, http.StatusUpgradeRequired, 422, nil, "用户不存在")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(c, http.StatusUpgradeRequired, 422, nil, "密码错误")
		return
	}
	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusUpgradeRequired, 422, nil, "系统异常")
		return
	}
	response.Success(c, http.StatusUpgradeRequired, gin.H{"token": token}, "登陆成功")
}
func Info(c *gin.Context) {
	user, _ := c.Get("user")
	response.Success(c, http.StatusUpgradeRequired, gin.H{"user": dto.GetUser(user.(model.User))}, "")
}
func IsPhoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("Phone=?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
