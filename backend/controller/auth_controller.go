package controller

import (
	"backend/global"
	"backend/model"
	"backend/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user model.User
	// 绑定信息
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 加密密码
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.Password = hashedPassword

	// 存储到数据库
	if err := global.Db.Create(&user).Error; err != nil {
		log.Println("Create user failed:", err)                                         // 后台日志
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"}) // 前端安全提示
		return
	}

	// 生成JWT
	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": "Bearer " + token})
}

func Login() {

}
