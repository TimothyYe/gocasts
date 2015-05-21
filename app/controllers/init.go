package controllers

import (
	"fmt"
	"github.com/jgraham909/revmgo"
	"github.com/revel/revel"
)

func init() {
	revmgo.ControllerInit()
	revel.InterceptFunc(preChecking, revel.BEFORE, &Admin{})
}

func preChecking(c *revel.Controller) revel.Result {
	if c.Session["user"] == "" {
		//User is not login
		fmt.Println("User is not login")
		return c.Redirect(App.Index)
	}

	return nil
}
