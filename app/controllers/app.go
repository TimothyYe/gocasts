package controllers

import (
	"code.google.com/p/go.crypto/bcrypt"
	"github.com/dchest/captcha"
	"github.com/jgraham909/revmgo"
	"github.com/revel/revel"
	"github.com/timothyye/gocasts/app/models"
	"labix.org/v2/mgo/bson"
)

type App struct {
	*revel.Controller
	revmgo.MongoController
}

func (c App) ShowCast(id string) revel.Result {
	t := models.Casts{}

	c.MongoSession.DB("gocasts").C("casts").FindId(bson.ObjectIdHex(id)).One(&t)
	cast := models.CastsView{Id: t.Id.Hex(), Author: t.Author, AuthorUrl: t.AuthorUrl,
		VisitCount: t.VisitCount, Title: t.Title, Intro: t.Intro,
		ShowNotes: t.ShowNotes, Url: t.Url, LogoUrl: t.LogoUrl, Date: t.Date, Tags: t.Tags}

	return c.Render(cast)
}

func (c App) SearchTag(tag string) revel.Result {
	casts := []models.Casts{}
	viewCasts := []models.CastsView{}

	c.MongoSession.DB("gocasts").C("casts").Find(bson.M{"tags": bson.RegEx{tag, ""}}).Sort("-date").All(&casts)

	for _, t := range casts {
		viewCasts = append(viewCasts,
			models.CastsView{Id: t.Id.Hex(), Author: t.Author, AuthorUrl: t.AuthorUrl,
				VisitCount: t.VisitCount, Title: t.Title, Intro: t.Intro,
				ShowNotes: t.ShowNotes, Url: t.Url, LogoUrl: t.LogoUrl, Date: t.Date, Tags: t.Tags})
	}

	num := len(viewCasts)
	pers := 12
	pager := NewPaginator(c.Request.Request, pers, num)

	return c.Render(viewCasts, pager, num, tag)
}

func (c App) Index() revel.Result {
	num, _ := c.MongoSession.DB("gocasts").C("casts").Count()

	pers := 9
	pager := NewPaginator(c.Request.Request, pers, num)

	casts := []models.Casts{}
	viewCasts := []models.CastsView{}

	c.MongoSession.DB("gocasts").C("casts").Find(nil).Limit(pers).Skip(pager.Offset()).Sort("-date").All(&casts)

	for _, t := range casts {
		viewCasts = append(viewCasts,
			models.CastsView{Id: t.Id.Hex(), Author: t.Author, AuthorUrl: t.AuthorUrl,
				VisitCount: t.VisitCount, Title: t.Title, Intro: t.Intro,
				ShowNotes: t.ShowNotes, Url: t.Url, LogoUrl: t.LogoUrl, Date: t.Date})
	}

	return c.Render(viewCasts, pager, num)
}

func (c App) CastsList() revel.Result {
	return nil
}

func (c App) About() revel.Result {
	return c.Render()
}

func (c App) LoginView() revel.Result {
	if c.Session["user"] != "" {
		return c.Redirect(Admin.Index)
	}
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

	id := models.Identity{}
	c.MongoSession.DB("gocasts").C("id").Find(bson.M{"email": "admin@gocasts.net"}).One(&id)

	if username == "admin@gocasts.net" && bcrypt.CompareHashAndPassword(id.Password, []byte(password)) == nil {
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
