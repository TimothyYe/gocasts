package controllers

import (
	"github.com/revel/revel"
	"github.com/timothyye/gocasts/app/models"
	"labix.org/v2/mgo/bson"
	"time"
)

func (c Admin) Casts() revel.Result {

	return c.Render()
}

func (c Admin) AddCastPage() revel.Result {
	return c.Render()
}

func (c Admin) AddCast(author, authorurl, title, intro, logourl, url, shownotes string) revel.Result {
	cast := models.Casts{Id: bson.NewObjectId(), Author: author, AuthorUrl: authorurl,
		VisitCount: 0, Title: title, Intro: intro, ShowNotes: shownotes,
		Url: url, LogoUrl: logourl, Date: time.Now().Format("2006-01-02 15:04:05")}
	c.MongoSession.DB("gocasts").C("casts").Insert(cast)

	return c.Redirect(Admin.Casts)
}

func (c Admin) ModifyCastPage() revel.Result {

	return c.Render()
}

func (c Admin) ModifyCast() revel.Result {

	return c.Render()
}

func (c Admin) RemoveCast() revel.Result {

	return c.Render()
}
