// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tAdmin struct {}
var Admin tAdmin


func (_ tAdmin) Casts(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Admin.Casts", args).Url
}

func (_ tAdmin) AddCastPage(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Admin.AddCastPage", args).Url
}

func (_ tAdmin) AddCast(
		author string,
		authorurl string,
		title string,
		tags string,
		intro string,
		logourl string,
		url string,
		shownotes string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "author", author)
	revel.Unbind(args, "authorurl", authorurl)
	revel.Unbind(args, "title", title)
	revel.Unbind(args, "tags", tags)
	revel.Unbind(args, "intro", intro)
	revel.Unbind(args, "logourl", logourl)
	revel.Unbind(args, "url", url)
	revel.Unbind(args, "shownotes", shownotes)
	return revel.MainRouter.Reverse("Admin.AddCast", args).Url
}

func (_ tAdmin) ModifyCastPage(
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Admin.ModifyCastPage", args).Url
}

func (_ tAdmin) ModifyCast(
		id string,
		author string,
		authorurl string,
		title string,
		tags string,
		intro string,
		logourl string,
		url string,
		shownotes string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "author", author)
	revel.Unbind(args, "authorurl", authorurl)
	revel.Unbind(args, "title", title)
	revel.Unbind(args, "tags", tags)
	revel.Unbind(args, "intro", intro)
	revel.Unbind(args, "logourl", logourl)
	revel.Unbind(args, "url", url)
	revel.Unbind(args, "shownotes", shownotes)
	return revel.MainRouter.Reverse("Admin.ModifyCast", args).Url
}

func (_ tAdmin) RemoveCast(
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Admin.RemoveCast", args).Url
}

func (_ tAdmin) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Admin.Index", args).Url
}

func (_ tAdmin) Password(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Admin.Password", args).Url
}

func (_ tAdmin) UpdatePassword(
		password string,
		verify string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "password", password)
	revel.Unbind(args, "verify", verify)
	return revel.MainRouter.Reverse("Admin.UpdatePassword", args).Url
}


type tApp struct {}
var App tApp


func (_ tApp) ShowCast(
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("App.ShowCast", args).Url
}

func (_ tApp) SearchTag(
		tag string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "tag", tag)
	return revel.MainRouter.Reverse("App.SearchTag", args).Url
}

func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}

func (_ tApp) CastsList(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.CastsList", args).Url
}

func (_ tApp) About(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.About", args).Url
}

func (_ tApp) LoginView(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.LoginView", args).Url
}

func (_ tApp) Login(
		username string,
		password string,
		captcha_id string,
		captcha_value string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "password", password)
	revel.Unbind(args, "captcha_id", captcha_id)
	revel.Unbind(args, "captcha_value", captcha_value)
	return revel.MainRouter.Reverse("App.Login", args).Url
}

func (_ tApp) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Logout", args).Url
}


type tCaptcha struct {}
var Captcha tCaptcha


func (_ tCaptcha) GetCaptchaImage(
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Captcha.GetCaptchaImage", args).Url
}

func (_ tCaptcha) GetCaptcha(
		captchaid string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "captchaid", captchaid)
	return revel.MainRouter.Reverse("Captcha.GetCaptcha", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


