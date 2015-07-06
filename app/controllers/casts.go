package controllers

import (
	"time"

	"github.com/revel/revel"
	"github.com/timothyye/gocasts/app/models"
	"labix.org/v2/mgo/bson"
)

func (c Admin) Casts() revel.Result {
	num, _ := c.MongoSession.DB("gocasts").C("casts").Count()

	pers := 12
	pager := NewPaginator(c.Request.Request, pers, num)

	casts := []models.Casts{}
	viewCasts := []models.CastsView{}

	c.MongoSession.DB("gocasts").C("casts").Find(nil).Limit(pers).Skip(pager.Offset()).Sort("-date").All(&casts)

	for _, t := range casts {
		viewCasts = append(viewCasts,
			models.CastsView{Id: t.Id.Hex(), Author: t.Author, AuthorUrl: t.AuthorUrl,
				VisitCount: t.VisitCount, Title: t.Title, Intro: t.Intro,
				ShowNotes: t.ShowNotes, Url: t.Url, LogoUrl: t.LogoUrl, Date: t.Date, Tags: t.Tags})
	}

	return c.Render(viewCasts, pager, num)
}

func (c Admin) AddCastPage() revel.Result {
	return c.Render()
}

func (c Admin) AddCast(author, authorurl, title, tags, intro, logourl, url, shownotes string) revel.Result {
	cast := models.Casts{Id: bson.NewObjectId(), Author: author, AuthorUrl: authorurl,
		VisitCount: 0, Title: title, Intro: intro, ShowNotes: shownotes,
		Url: url, LogoUrl: logourl, Date: time.Now().Format("2006-01-02"), Tags: tags}
	c.MongoSession.DB("gocasts").C("casts").Insert(cast)

	return c.Redirect(Admin.Casts)
}

func (c Admin) ModifyCastPage(id string) revel.Result {
	t := models.Casts{}

	c.MongoSession.DB("gocasts").C("casts").FindId(bson.ObjectIdHex(id)).One(&t)
	cast := models.CastsView{Id: t.Id.Hex(), Author: t.Author, AuthorUrl: t.AuthorUrl,
		VisitCount: t.VisitCount, Title: t.Title, Intro: t.Intro,
		ShowNotes: t.ShowNotes, Url: t.Url, LogoUrl: t.LogoUrl, Date: t.Date, Tags: t.Tags}

	return c.Render(cast)
}

func (c Admin) ModifyCast(id, author, authorurl, title, tags, intro, logourl, url, shownotes string) revel.Result {
	c.MongoSession.DB("gocasts").C("casts").UpdateId(bson.ObjectIdHex(id),
		bson.M{"$set": bson.M{"author": author,
			"authorurl": authorurl,
			"title":     title,
			"intro":     intro,
			"shownotes": shownotes,
			"url":       url,
			"logourl":   logourl,
			"tags":      tags}})

	return c.Redirect(Admin.Casts)
}

func (c Admin) RemoveCast(id string) revel.Result {
	c.MongoSession.DB("gocasts").C("casts").RemoveId(bson.ObjectIdHex(id))
	return c.Redirect(Admin.Casts)
}
