package tools

import (
	"bytes"
	"github.com/golang/freetype/truetype"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"io"
	"sync"
)

const (
	CaptchaHeight = 120
	CaptchaWight  = 240
	NoiseCnt      = 2
	LineOpt       = 3
)

var (
	mathCaptcha     *base64Captcha.DriverMath
	mathCaptchaOnce = &sync.Once{}
	langCaptcha     *base64Captcha.DriverLanguage
	langCaptchaOnce = &sync.Once{}
)

// 生成数学字符串验证码图片
func MathCaptcha() (id, aws string, img string, err error) {
	mathCaptchaOnce.Do(func() {
		mathCaptcha = base64Captcha.NewDriverMath(CaptchaHeight, CaptchaWight, NoiseCnt, LineOpt, &color.RGBA{
			R: uint8(RandIntN(255)),
			G: uint8(RandIntN(255)),
			B: uint8(RandIntN(255)),
			A: uint8(RandIntN(255)),
		}, []string{"RitaSmith.ttf"})
	})
	id, ques, aws := mathCaptcha.GenerateIdQuestionAnswer()
	item, err := mathCaptcha.DrawCaptcha(ques)
	if err != nil {
		return "", "", "", err
	}

	return id, aws, item.EncodeB64string(), nil
}

// 生成中文验证码
func CharCaptcha() (id, aws string, img io.Reader, err error) {
	langCaptchaOnce.Do(func() {
		langCaptcha = base64Captcha.NewDriverLanguage(CaptchaHeight, CaptchaWight, NoiseCnt, LineOpt, 4, &color.RGBA{
			R: uint8(RandIntN(255)),
			G: uint8(RandIntN(255)),
			B: uint8(RandIntN(255)),
			A: uint8(RandIntN(255)),
		}, []*truetype.Font{}, "zh")
	})
	id, ques, aws := langCaptcha.GenerateIdQuestionAnswer()
	item, err := langCaptcha.DrawCaptcha(ques)
	if err != nil {
		return "", "", nil, err
	}
	buf := bytes.NewBuffer(make([]byte, 0))

	_, err = item.WriteTo(buf)
	if err != nil {
		return "", "", nil, err
	}

	return id, aws, buf, nil
}
