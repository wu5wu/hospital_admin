package ruoters

import (
	"fmt"
	"gin_vue/controllers"
	"gin_vue/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"strconv"
)

func IndexRuoteraInit(r *gin.Engine) {
	r.GET("/ping", controllers.IndexController{}.User)

	r.GET("/news", controllers.IndexController{}.News)

	r.GET("/", controllers.IndexController{}.Index)

	r.POST("/index", func(c *gin.Context) {
		username := c.PostForm("username")
		pwd := c.PostForm("pwd")
		c.HTML(http.StatusOK, "news.html", gin.H{
			"title": "后端数据",
		})
		c.JSON(200, gin.H{
			"username": username,
			"pwd":      pwd,
		})
	})

	//单文件上传
	r.POST("/admin/user/doUpload", func(c *gin.Context) {
		username := c.PostForm("username")
		file, err := c.FormFile("face")
		if err == nil {

			allowExtMap := map[string]bool{
				".jpg":  true,
				".png":  true,
				".gif":  true,
				".jpeg": true,
			}
			if _, ok := allowExtMap[path.Ext(file.Filename)]; !ok {
				c.String(200, "文件不合法")
			}
			//创建文件保存目录
			day := models.GetDay()
			dir := "./static/images/" + day
			err := os.MkdirAll(dir, 0666)
			if err != nil {
				fmt.Println("创建目录失败")
				return
			}

			//创建文件名进行保存
			unix := models.GetUnix()
			fileName := strconv.FormatInt(unix, 10) + path.Ext(file.Filename)
			dst := path.Join(dir, fileName)
			c.SaveUploadedFile(file, dst)

			c.JSON(200, gin.H{
				"username": username,
				"path":     dst,
			})
		}

	})

	//duo文件上传
	r.POST("/admin/user/doEdit", func(c *gin.Context) {
		username := c.PostForm("username")
		form, _ := c.MultipartForm()
		file := form.File["face[]"]
		for _, file := range file {
			dst := path.Join("./static/images", file.Filename)
			c.SaveUploadedFile(file, dst)
		}

		c.JSON(200, gin.H{
			"username": username,
		})
	})
	r.POST("/alipay/sandbox/tradePagePay", controllers.AlipayController{}.TradePagesandBoxPay)
}
