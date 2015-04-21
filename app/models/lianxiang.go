package models


import (
	//"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//"github.com/revel/revel"
)

type Lianxiang struct {
	Id_             bson.ObjectId `bson:"_id"`
	Type string
	Content string  `bson:"content"`
	Oid string `bson:"oid"`
}

