package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"net/http"
)

// YanzhengController 控制器
type YanzhengController struct{}

// CaptchaData 结构体用于表示验证码的数据
type CaptchaData struct {
	CaptchaId string `json:"captcha_id"` // 验证码id
	Data      string `json:"data"`       // 验证码数据base64类型
	Answer    string `json:"answer"`     // 验证码数字
}

// 数字驱动配置
var digitDriver = base64Captcha.DriverDigit{
	Height:   50,
	Width:    200,
	Length:   4,   // 验证码长度
	MaxSkew:  0.7, // 倾斜
	DotCount: 1,   // 背景的点数，越大，字体越模糊
}

// 存储配置
var store = base64Captcha.DefaultMemStore

// CaptchaGenerate 生成验证码
func CaptchaGenerate() (CaptchaData, error) {
	var ret CaptchaData
	c := base64Captcha.NewCaptcha(&digitDriver, store)
	id, b64s, answer, err := c.Generate()
	if err != nil {
		return ret, err
	}
	ret.CaptchaId = id
	ret.Data = b64s
	ret.Answer = answer
	return ret, nil
}

// CaptchaVerify 验证验证码
func CaptchaVerify(data CaptchaData) bool {
	return store.Verify(data.CaptchaId, data.Answer, true)
}

// GenCaptcha 生成验证码的处理函数
func (con YanzhengController) GenCaptcha(c *gin.Context) {
	captcha, err := CaptchaGenerate()
	if err != nil {
		//zap.L().Error("GenCaptcha CaptchaGenerate() failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate captcha"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": captcha, "code": 200})
}

// CaptchaVerify 处理验证码验证的函数
func (con YanzhengController) VerifyCaptcha(c *gin.Context) {
	var param CaptchaData
	if err := c.ShouldBindJSON(&param); err != nil {
		//zap.L().Error("CaptchaVerify c.ShouldBindJSON(&param) failed", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	fmt.Println("param:", param)
	if !CaptchaVerify(param) {
		c.JSON(http.StatusOK, gin.H{"message": "验证失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "验证成功"})
}
