package models

import (
	"labix.org/v2/mgo/bson"
)

type Casts struct {
	Id         bson.ObjectId `bson:"_id,omitempty"`
	Author     string
	AuthorUrl  string
	VisitCount int
	Title      string
	LogoUrl    string
	Intro      string
	ShowNotes  string
	Url        string
	Date       string
	Tags       string
}

type CastsView struct {
	Id         string
	Author     string
	AuthorUrl  string
	VisitCount int
	Title      string
	LogoUrl    string
	Intro      string
	ShowNotes  string
	Url        string
	Date       string
	Tags       string
}
