package models

//用户相关

import (
	//"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/revel/revel"
)

// 用户
type User struct {
	Id_             bson.ObjectId `bson:"_id"`
	Username        string
	Password        string
	Salt            string `bson:"salt"`
	Email           string
	IsSuperuser     bool
	IsActive        bool
}

func Login(s revel.Session, u *User){
	s["username"] = u.Username
	s["email"] = u.Email
}

func Logout(s *revel.Session){
	*s = make(revel.Session)
}

func IsLogin(s revel.Session) bool {
	if _, found := s["username"];found {
		return true
	}
	return false
}
