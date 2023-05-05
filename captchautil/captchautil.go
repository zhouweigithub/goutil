package captchautil

import (
	"github.com/mojocn/base64Captcha"
)

var (
	conf base64Captcha.DriverString
	cap  *base64Captcha.Captcha
)

// 初始化
func init() {
	conf = base64Captcha.DriverString{
		Width:           240,
		Height:          80,
		Length:          5,
		NoiseCount:      80,
		Source:          base64Captcha.TxtNumbers,
		ShowLineOptions: base64Captcha.OptionShowHollowLine,
		// Fonts:           []string{""},
	}
	// loads fonts by names
	var driver = conf.ConvertFonts()
	cap = base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
}

// 生成验证码并存入缓存
//
//	id: 验证码缓存id
//	base64String: 验证码图片内容
//	err: 错误信息
func GetImgCaptcha() (id, base64String string, err error) {
	id, base64String, err = cap.Generate()
	return id, base64String, err
}

// 验证验证码是否正确
//
//	id: 验证码缓存id
//	answer: 验证码
func Verify(id, answer string) bool {
	return cap.Verify(id, answer, true)
}
