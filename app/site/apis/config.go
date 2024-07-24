package apis

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/google/uuid"
	"github.com/mojocn/base64Captcha"
)

type Config struct {
	api.Api
}
type configJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

func DriverDigitFunc() (id, b64s string, err error) {
	e := configJsonBody{}
	e.Id = uuid.New().String()

	e.DriverDigit = base64Captcha.NewDriverDigit(80, 240, 4, 0.2, 16)
	driver := e.DriverDigit

	return base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore).Generate()
}

func (e Config) Captcha(c *gin.Context) {

	err := e.MakeContext(c).Errors
	if err != nil {
		e.Error(500, err, "服务初始化失败！")
		return
	}
	id, b64s, err := DriverDigitFunc()
	if err != nil {
		e.Logger.Errorf("DriverDigitFunc error, %s", err.Error())
		e.Error(500, err, "验证码获取失败")

		return
	}

	data := map[string]interface{}{
		"base64": b64s,
		"cid":    id,
	}
	e.OK(data, "success")
}
func (e Config) Config(c *gin.Context) {
	err := e.MakeContext(c).
		MakeOrm().
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	dat := `{
  "loginRegisterSwitch": 1,
  "loginCaptchaSwitch": 1,
  "loginAvatar": "http://bufanyun.cn-bj.ufileos.com/hotgo/attachment/2023-02-09/cqdq8er9nfkchdopav.png",
  "loginRoleId": 202,
  "loginDeptId": 109,
  "loginPostIds": [
    6
  ],
  "loginProtocol": "<p><span style=\"color: rgb(31, 34, 37);\">用户协议..</span></p>",
  "loginPolicy": "<p><span style=\"color: rgb(31, 34, 37);\">隐私权政策..</span></p>",
  "loginAutoOpenId": 2,
  "loginForceInvite": 2
}`

	data := make(map[string]interface{}, 0)

	err = json.Unmarshal([]byte(dat), &data)
	if err != nil {

		e.Logger.Error(err)
		e.Error(-1, nil, err.Error())
		return
	}
	e.OK(data, "successful")
	return
}
