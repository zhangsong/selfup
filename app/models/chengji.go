package models

//成绩

import (
	//"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//"github.com/revel/revel"
)

// 用户
type Chengji struct {
	Id_             bson.ObjectId `bson:"_id"`
	ShijuanId string
	Content string //答题详情
	Count int //分数
}

