package controllers

import (
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
	return c.Render()
}

func (c App) Login(username, password string) revel.Result {

	if username == "" || password == "" {
		return c.Redirect(App.Login)
	}

	if username == "admin@abc.com" && password == "admin888" {
		c.Session["user"] = "admin"
		return c.Redirect(Admin.Index)
	}

	return c.Redirect(App.Login)
}

func (c App) Logout() revel.Result {
	//Clear session
	c.Session["user"] = ""
	return c.Redirect(App.Index)
}
