package controllers

import (
	"code.google.com/p/go.crypto/bcrypt"
	"github.com/jgraham909/revmgo"
	"github.com/revel/revel"
	"labix.org/v2/mgo/bson"
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

func (c Admin) UpdatePassword(password, verify string) revel.Result {
	var isSuccess bool
	isPost := true

	if password == verify {

		hashedPass, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		info, err := c.MongoSession.DB("gocasts").C("id").Upsert(bson.M{"email": "admin@gocasts.net"}, bson.M{"$set": bson.M{"password": hashedPass, "email": "admin@gocasts.net"}})

		if err == nil {
			if info.Updated >= 0 {
				isSuccess = true
			}
		}
	} else {
		isSuccess = false
	}
	c.RenderArgs["isPost"] = isPost
	c.RenderArgs["isSuccess"] = isSuccess

	return c.RenderTemplate("admin/Password.html")
}
