package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var db *gorm.DB

func init() {
	var err error
	dsn := "root:你的密码@tcp(127.0.0.1:3306)/crud_demo"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败: ", err)
	}
	db.AutoMigrate(&User{})
	log.Println("数据库连接并初始化成功！")
}

func main() {
	r := gin.Default()

	r.POST("/users", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "数据格式错误"})
			return
		}
		db.Create(&user)
		c.JSON(http.StatusOK, gin.H{"message": "创建成功", "user": user})
	})

	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user User
		result := db.First(&user, id)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
			return
		}
		c.JSON(http.StatusOK, user)
	})

	r.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user User
		if err := db.First(&user, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
			return
		}
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "数据格式错误"})
			return
		}
		db.Save(&user)
		c.JSON(http.StatusOK, gin.H{"message": "更新成功", "user": user})
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		db.Delete(&User{}, id)
		c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
	})

	r.Run(":8080")
}
