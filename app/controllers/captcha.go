package controllers

import (
	"fmt"
	"github.com/dchest/captcha"
	"github.com/revel/revel"
)

type Captcha struct {
	*revel.Controller
}

type CaptchaResult struct {
	ID       string `json:"id"`
	ImageURL string `json:"imageurl"`
}

func (c Captcha) GetCaptchaImage(id string) revel.Result {
	c.Response.ContentType = "image/png"
	captcha.WriteImage(c.Response.Out, id, 300, 60)
	return nil
}

func (c Captcha) GetCaptcha(captchaid string) revel.Result {
	capResult := new(CaptchaResult)

	capResult.ID = captcha.NewLen(4)

	capResult.ImageURL = fmt.Sprintf(
		"/captcha/img?id=%s",
		capResult.ID,
	)
	return c.RenderJson(capResult)
}
