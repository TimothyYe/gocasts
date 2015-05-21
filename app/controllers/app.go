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

func (c App) Login() revel.Result {
	return c.Render()
}

func (c App) Logout() revel.Result {
	return c.Render()
}
