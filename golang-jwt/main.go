package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func generateToken(c *gin.Context, userId uint) (string, error) {
	claims := createClaims(userId)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			c.Abort()
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// 可以将claims中的数据附加到Context中，供后续处理使用
			c.Set("userId", claims["userId"])
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			c.Abort()
		}
	}
}

var jwtSecret = []byte("your-secret-key")

func createClaims(userId uint) jwt.MapClaims {
	return jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // 设置过期时间为24小时后
	}
}

func main() {
	r := gin.Default()

	r.POST("/login", func(c *gin.Context) {
		account := c.Param("account")
		password := c.Param("password")
		fmt.Printf("account=%s\n", account)
		fmt.Printf("password=%s\n", password)
		//账号密码正确查到id=999
		token, err := generateToken(c, 999)
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "Welcome!", "token": token})
	})

	// 示例路由，使用JWT验证
	protected := r.Group("/api/v1", authMiddleware())
	{
		protected.GET("/profile", func(c *gin.Context) {
			userId := c.MustGet("userId").(uint)
			c.JSON(http.StatusOK, gin.H{"message": "Welcome!", "userId": userId})
		})
	}

	r.Run(":8080") // 启动服务
}
