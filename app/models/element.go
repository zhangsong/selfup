package models

//题元

import (
	//"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//"github.com/revel/revel"
)

// 用户
type Element struct {
	Id_             bson.ObjectId `bson:"_id"`
	Yiyong	int
	Content string
	Hash	string
	Aanda	string
}

