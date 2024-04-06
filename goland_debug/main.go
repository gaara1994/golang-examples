/**
 * @author yantao
 * @date 2024/4/6
 * @description Goland的debug使用
 */
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Spec map[string]interface{} // 自定义Spec类型，用于表示jsonb中的任意结构

type Goods struct {
	ID    uint    `gorm:"primaryKey"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Spec  Spec    `gorm:"type:jsonb;column:spec" json:"spec"` // 使用map[string]interface{}映射jsonb字段
}

var db *gorm.DB

func main() {
	// 1. 连接数据库
	dsn := "user=postgres password=MyNewPass4! host=localhost port=5432 dbname=mydatabase sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移schema
	db.AutoMigrate(&Goods{})

	// 2. 设置Gin路由和处理器
	r := gin.Default()

	r.GET("/goods", GetGoods)
	r.POST("/goods", CreateGoods)
	r.PUT("/goods/:id", UpdateGoods)
	r.DELETE("/goods/:id", DeleteGoods)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}

// GetGoods 处理GET请求，返回所有商品
func GetGoods(c *gin.Context) {
	var goods []Goods
	db.Find(&goods)
	c.JSON(http.StatusOK, goods)
}

// CreateGoods 处理POST请求，创建新商品
func CreateGoods(c *gin.Context) {
	var good Goods
	c.BindJSON(&good)
	db.Create(&good)
	c.JSON(http.StatusOK, good)
}

// UpdateGoods 处理PUT请求，更新指定ID的商品
func UpdateGoods(c *gin.Context) {
	var good Goods
	c.BindJSON(&good)
	id := c.Param("id")
	fmt.Println("ID: ", id)
	db.Model(&Goods{}).Where("id = ?", id).Updates(&good)
	c.Status(http.StatusOK)
}

// DeleteGoods 处理DELETE请求，删除指定ID的商品
func DeleteGoods(c *gin.Context) {
	id := c.Param("id")
	db.Delete(&Goods{}, id)
	c.Status(http.StatusOK)
}
