package tools

import (
	"bytes"
	"github.com/golang/freetype/truetype"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"io"
)

const (
	MaxColor = 2 << 7
	LineOpt  = 3
)

type CaptchaOpts struct {
	Height   int
	Width    int
	NoiseCnt int
}

// 生成数学字符串验证码图片
func MathCaptcha(opt *CaptchaOpts) (id, aws string, img string, err error) {
	driver := base64Captcha.NewDriverMath(opt.Height, opt.Width, opt.NoiseCnt, LineOpt, &color.RGBA{
		R: uint8(RandIntN(MaxColor)),
		G: uint8(RandIntN(MaxColor)),
		B: uint8(RandIntN(MaxColor)),
		A: uint8(RandIntN(MaxColor)),
	}, []string{"RitaSmith.ttf"})

	id, ques, aws := driver.GenerateIdQuestionAnswer()
	item, err := driver.DrawCaptcha(ques)
	if err != nil {
		return "", "", "", err
	}

	return id, aws, item.EncodeB64string(), nil
}

// 生成中文验证码
func CharCaptcha(opt *CaptchaOpts) (id, aws string, img io.Reader, err error) {
	driver := base64Captcha.NewDriverLanguage(opt.Height, opt.Width, opt.NoiseCnt, LineOpt, 4, &color.RGBA{
		R: uint8(RandIntN(MaxColor)),
		G: uint8(RandIntN(MaxColor)),
		B: uint8(RandIntN(MaxColor)),
		A: uint8(RandIntN(MaxColor)),
	}, []*truetype.Font{}, "zh")

	id, ques, aws := driver.GenerateIdQuestionAnswer()
	item, err := driver.DrawCaptcha(ques)
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
