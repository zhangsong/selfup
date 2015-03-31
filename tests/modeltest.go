package tests

import (
	"github.com/revel/revel"
	"github.com/zhangsong/selfup/app/models"
	"reflect"
	"gopkg.in/mgo.v2"
	)

type ModelTest struct {
	revel.TestSuite
}

func (t *ModelTest) Before() {
	println("Set up")
}

func (t *ModelTest) TestMgoInstanceCreateOk() {
	m := models.NewMyMgo("s")
	t.Assert(m!=nil)
	
	if url, found := revel.Config.String("mgo.url"); found {
		m = models.NewMyMgo(url)
		s := m.GetInstance()
		t.AssertEqual(reflect.TypeOf(s),reflect.TypeOf((*mgo.Session)(nil)))
		
		if testdb,f:=revel.Config.String("mgo.testdb");f{		
			db := m.DB(testdb)
			t.AssertEqual(reflect.TypeOf(db), reflect.TypeOf((*mgo.Database)(nil)))
		} else {
			panic("请设定测试数据库")
		}
		
	} else {
		panic("mgo配置项目错误")
	}
}

func (t *ModelTest) After() {
	println("Tear down")
}
