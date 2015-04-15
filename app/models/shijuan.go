package models

//试卷

import (
	//"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//"github.com/revel/revel"
)

// 用户
type Shijuan struct {
	Id_             bson.ObjectId `bson:"_id"`
	Content string
}

type Juger interface {
	Juge(eid string, daan []string)([]string, bool)
}

