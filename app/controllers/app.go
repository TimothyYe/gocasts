package controllers

import (
	"github.com/dchest/captcha"
	"github.com/jgraham909/revmgo"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
	revmgo.MongoController
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) About() revel.Result {
	return c.Render()
}

func (c App) LoginView() revel.Result {

	CaptchaId := captcha.NewLen(4)
	return c.Render(CaptchaId)
}

func (c App) Login(username, password, captcha_id, captcha_value string) revel.Result {

	c.Validation.Required(username).Message("请输入用户名")
	c.Validation.Required(password).Message("请输入密码")

	if !captcha.VerifyString(captcha_id, captcha_value) {
		c.Validation.Error("验证码不正确")
	}

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()

		return c.Redirect(App.LoginView)
	}

	if username == "admin@abc.com" && password == "admin888" {
		c.Session["user"] = "admin"
		return c.Redirect(Admin.Index)
	}

	return c.Redirect(App.LoginView)
}

func (c App) Logout() revel.Result {
	//Clear session
	c.Session["user"] = ""
	return c.Redirect(App.Index)
}
