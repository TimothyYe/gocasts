package controllers

import (
	"github.com/jgraham909/revmgo"
	"github.com/revel/revel"
)

type Admin struct {
	*revel.Controller
	revmgo.MongoController
}

func (c Admin) Index() revel.Result {
	return c.Render()
}

func (c Admin) Password() revel.Result {
	return c.Render()
}
