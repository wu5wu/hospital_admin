package controllers

import (
	"fmt"
	"gin_vue/models"
	"github.com/gin-gonic/gin"
	"os"
	"path"
	"strconv"
)

type ToolsController struct {
}

// 图片上传
func (con ToolsController) Login(c *gin.Context) {

	//username := c.PostForm("username")
	file, err := c.FormFile("file")
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
		imgurl := "http://localhost:8080/patient/getimg?imageName=" + dst
		c.JSON(200, gin.H{
			//"username": username,
			"path": imgurl,
		})
	}
}

// 显示图片的方法
func (con ToolsController) GetImage(c *gin.Context) {
	imageName := c.Query("imageName")  //截取get请求参数，也就是图片的路径，可是使用绝对路径，也可使用相对路径
	file, _ := os.ReadFile(imageName)  //把要显示的图片读取到变量中
	c.Writer.WriteString(string(file)) //关键一步，写给前端
}
