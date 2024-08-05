package controllers

import (
	"fmt"
	"gin_vue/models"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/v3"
	"net/http"
	"strconv"
	"time"
)

type AlipayController struct {
}
type AlipaySandBox struct {
	AppID        string
	PrivateKey   string
	AliPublicKey string
	NotifyURL    string
	ReturnURL    string
}

var PatientId int
var Balance int

// 调用支付宝沙箱
func (con AlipayController) TradePagesandBoxPay(c *gin.Context) {
	//var deals deal
	//var deals2 deal
	var updateData models.Patient
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}
	PatientId = updateData.PatientId
	Balance = updateData.Balance
	orderNo := fmt.Sprintf("%d", time.Now().Unix())
	p := alipay.TradePagePay{}

	p.NotifyURL = "http://2y7is8.natappfree.cc/patient/topup"
	p.ReturnURL = "http://localhost:5173/patient/information"
	p.Subject = "充值"
	p.OutTradeNo = orderNo
	p.TotalAmount = strconv.Itoa(updateData.Balance)
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"
	AlipayClient, err := alipay.New("9021000137684221", "MIIEogIBAAKCAQEAnf2gADwVp1Ki8vapa4KU4TFw+Mfkbarfg5UcCkIX9BnX4IO9ApseXe6NBC5832xm1cPXgfgxH6BdMi9cMazARGtFYj7xYJQLNtgCXUNDHYuhe2XcFta7hZtOvkTVDmV7SuRgGU0BgznwHKZxjiAxEtPCdhxQUTU+JFvOH7OxV/jcfia0s+or5/KPxgEe3A/bz3keRob+SPrbwqc4x771uL2kfAPPdoDxYYg7lAOONVqqDOIy29y28a0OtdZHJ99FZK8uVZWsmWPYWe9+Zd/r8sPlgJszHigwXp42r4mSsb23r4iLixhzr10RDjB7tfDZgJ48+kOh1jLXwHm+VpgtCQIDAQABAoIBAHNqSGxjeF4XnDppvzJu3lv0Rlc3j8Qdw9LYJvfHeMA18OJu2rEAqJpuHk4jt81v31/iBZlIunokKHD08CDJ1lslEjbHTIYFIHwcKP4AmRMIYhtHOhTlIXBlGI0es5YkxIDdfOZ+vtAijyoe3W26Tp82WyT1YicxAgiFmHrIRpHAbxQRRuTcKdQJJsxBZdo8pOLXWVAOnEyLEFqqM0wVXahaTirlBDyFBJdswCcSDB+1NYaQ5pnTlms3tdXD3BUP5DH9uZbzFDR9B9s1l2x4G9i8LxwGepnSuOLv4Lnp1AYr/64+Ld0ZbX7BmBuXPLU8dp0NnOEgBTf+Vy5ZNETDHFECgYEA49ASQVwe44fOWkxv8UikqGYgufVRHm3GQ5TiyZC0K5q2Jod+yZatj0Q5zOeMHf6LO9Oe62FIKHKxIuVZ21tCLI2OAubuPTnUKNKFFI1uuQigUTVCGErUKadc+aP428eTW6prrs2M14at9Zu1gNOib4eQKhuEMJYT3rda5Ji0ZjcCgYEAsYnz4ZxvYgpxXR8tiemPApAOr5u1EgnWAmKgUev8bbAiiPw5T97gvsR47wcOjZI9RpeCZAScTEiJ6hLz40rZZ/fSjCXX/d3KRGRqjJHMVX7VEAa0Ht3HuVgtsOtBPtxON9x8YhSqNARUCmpgq6MC3soKqZ+ljLadudmdSS/qZr8CgYBGJ49XnVzNyMadPCjE7w/14+10Fr2yVy/VJAZq+CrjurMZTmSXA9XS5rm+Bhg6LGcIN3UbG0wzWI6AJozzZjkMq1LuRdtb83EIWTwuRhew1503JOWDwZdwdF/HaN18IQuaEPM0U4xp2sm7dwX/9CxzmBWnQe4FYFIiJejRuMzlJQKBgDFGFV7fePZf6Zirx4rIPL28V6cbcwS/oBrnXpF+CoAy61LTleB5/ixUATBt5/cYyFLoR/MUaICvfT9u5SuHIZG+cTweYLT/slyK9htvd6Oe4OliwujyKrVhKMU6ddBJepf8ZLUgASEQtnZxkcTxKCHHN3Bec6b7JAlpwUKHvumnAoGANN39JWjBbjKJWSxYWc+ndFRSgH+4lkHOpgz7KJENUpfEvFDiD4oc/3jry7cCkqlaeh8mtruFN3WKdew4kSx3PMH30XDQaCZS1WCKcJUhiE881o2BMIlWBS9bSNR2ei6JDL/OSphxPNl7ureQ8nkMaM1olIOsCMHs31YVkD8O1uU=", false)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 201, "message": "初始化失败"})
		return
	}
	err = AlipayClient.LoadAlipayCertPublicKey("MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAv9+oqRhORPiNmZLT3o1bsdX/8uDnsXkiPyibG+Zs4qZQReKX+YfbO9w+kLyJPXlz7b2uthOAUimxHAEukYmV+fwfmlI3YF2jmy1dVBxrCi57V8SuKC31vW/CsBQDk00F9XaM+8zVO8hfU42mTymtCH/OAJV0lyXwo4s2Sfo48uQEV2Q9BwnbZGsnlRQ2VC7ynVdJafy3vuMVsFbt/OPan/hF3by16ULFzhr74H8lGKx4Hi5nxamAt1OVrBcsimqOfNN34ytL4uQzb0XvkDibuYM+VahSjhJoq3aeZCOHLhL/UsOSfU50ikUBezPDmjhccMIpVi46D3qYEYHrtz89DQIDAQAB")
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"code": 201, "message": err})
		return
	}
	url, err := AlipayClient.TradePagePay(p)
	if err != nil {
		fmt.Println("生成URL失败")
		return
	}
	var payUrl = url.String()
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "订单生成成功",
		"result":  payUrl,
	})

}

// 进行余额充值
func (con AlipayController) TopUp(c *gin.Context) {

	var deals deal

	models.DB.Raw("SELECT balance FROM patient WHERE patient_id = ?", PatientId).Scan(&deals) //获取数据库中的余额
	balance := deals.Balance + Balance
	var aaa int
	models.DB.Raw("UPDATE patient SET balance = ? WHERE patient_id = ?", balance, PatientId).Scan(aaa)
	c.JSON(200, gin.H{"message": "充值 successfully", "code": 200, "balance": balance})
	return

}

// 获取余额
func (con AlipayController) GetBalance(c *gin.Context) {

	var deals deal
	var updateData models.Patient
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}

	models.DB.Raw("SELECT balance FROM patient WHERE patient_id = ?", updateData.PatientId).Scan(&deals)
	c.JSON(200, gin.H{"code": 200, "balance": deals.Balance})
	return

}
