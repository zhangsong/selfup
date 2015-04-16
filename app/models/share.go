package models


import (
	//"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//"github.com/revel/revel"
)

// 分享
type Share struct {
	Id_             bson.ObjectId `bson:"_id"`
	ShijuanId string
}
