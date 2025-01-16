package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Mock 数据库，存储授权码及其状态
var mockDB = map[string]bool{
	"authcode123": true,
	"authcode456": false,
}

// validateAuthCode 处理 POST 请求，验证授权码
func validateAuthCode(c *gin.Context) {
	var requestBody struct {
		AuthCode string `json:"auth_code"`
	}

	// 解析请求体
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// 验证授权码
	if isValid, ok := mockDB[requestBody.AuthCode]; ok && isValid {
		c.JSON(http.StatusOK, gin.H{"status": "valid"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "invalid"})
	}
}

// Start 启动 HTTP 服务
func Start() {
	r := gin.Default()

	// 定义 POST 路由
	r.POST("/validate", validateAuthCode)

	// 启动服务，监听 8080 端口
	r.Run(":8080")
}