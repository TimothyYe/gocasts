package controllers

import (
	"fmt"
	"github.com/jgraham909/revmgo"
	"github.com/revel/revel"
	"strings"
)

func init() {
	revmgo.ControllerInit()
	revel.InterceptFunc(preChecking, revel.BEFORE, &Admin{})
	revel.TemplateFuncs["parseTags"] = parseTags
}

func parseTags(tags string) []string {
	return strings.Split(tags, ",")
}

func preChecking(c *revel.Controller) revel.Result {
	if c.Session["user"] == "" {
		//User is not login
		fmt.Println("User is not login")
		return c.Redirect(App.Index)
	}

	return nil
}
