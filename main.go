package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null"`
	Phone    string `gorm:"type:varchar(11);not null;unique"`
	Password string `gorm:"size:255;not null;"`
}

func main() {
	db := InitDB()
	defer db.Close()
	r := gin.Default()
	//localhost:8080
	r.POST("/api/auth/register", func(c *gin.Context) {
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
			name = RandomString(10)
		}
		if IsPhoneExist(db, phone) {
			c.JSON(http.StatusUpgradeRequired, gin.H{"code": 422, "msg": "用户已经存在"})
			return
		}
		newuser := User{
			Name:     name,
			Phone:    phone,
			Password: password,
		}
		db.Create(&newuser)

		log.Println(name, password, phone)
		c.JSON(200, gin.H{
			"message": "注册成功",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
func RandomString(n int) string {
	var letter = []byte("dsabkjdbsakjdhksahdkjahdkhsadkjshda")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letter[rand.Intn(len(letter))]
	}
	return string(result)
}
func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "ginessential"
	username := "root"
	password := ""
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}
	db.AutoMigrate(&User{})
	return db
}
func IsPhoneExist(db *gorm.DB, phone string) bool {
	var user User
	db.Where("Phone=?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
