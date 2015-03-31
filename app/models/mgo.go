package models

//mgo初始化

import (
	"gopkg.in/mgo.v2"
)

type myMgo func() *mgo.Session

func (m myMgo)GetInstance() *mgo.Session {
	return m()
}

func (m myMgo)DB(db string) *mgo.Database {
	return m.GetInstance().DB(db)
}


func NewMyMgo(url string) myMgo {
   var session *mgo.Session
   var err error
   return func()*mgo.Session{
	if nil==session {
		session, err = mgo.Dial(url)
		if err != nil {
			panic(err)
		}
		session.SetMode(mgo.Monotonic, true)
	}
	return session
   }

}