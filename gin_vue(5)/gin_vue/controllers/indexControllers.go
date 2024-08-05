package controllers

import (
	"gin_vue/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
}

func (con IndexController) User(c *gin.Context) {
	tx := models.DB.Begin() // 开启事务
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() // 回滚事务
		}
	}()

	var userList []models.Patient
	if err := tx.Raw("SELECT * FROM patient").Scan(&userList).Error; err != nil {
		// 查询失败，处理错误
		tx.Rollback() // 回滚事务
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
		return
	}

	// 查询成功，返回数据
	c.JSON(http.StatusOK, gin.H{"result": userList})
}

func (con IndexController) News(c *gin.Context) {
	c.HTML(http.StatusOK, "news.html", gin.H{
		"title": "后端数据",
	})
}

func (con IndexController) Index(c *gin.Context) { //路由传值
	username := c.Query("username") //获取路由中的key对应的值
	pwd := c.Query("pwd")
	age := c.DefaultQuery("age", "18")
	c.JSON(200, gin.H{
		"username": username,
		"pwd":      pwd,
		"age":      age,
	})
}
