package models

//试卷

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//"github.com/revel/revel"
)

// 用户
type Shijuan struct {
	Id_             bson.ObjectId `bson:"_id"`
	Content string
	Name string
}

type Juger interface {
	Juge(eid string, daan []string)([]string, bool)
}

func GetName(db *mgo.Database, id string) (name string) {
	var e Shijuan
	db.C(SJ).Find(bson.M{"_id":bson.ObjectIdHex(id)}).One(&e)
	return e.Name
}

