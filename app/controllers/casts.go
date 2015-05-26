package controllers

import (
	"fmt"
	"github.com/revel/revel"
)

func (c Admin) Casts() revel.Result {

	return c.Render()
}

func (c Admin) AddCastPage() revel.Result {
	return c.Render()
}

func (c Admin) AddCast(author, authorurl, title, intro, logourl, url, shownotes string) revel.Result {
	fmt.Println("author:" + author)
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
