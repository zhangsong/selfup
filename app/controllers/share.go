package controllers

import (
	"github.com/revel/revel"
	"github.com/zhangsong/selfup/app/models"
	"gopkg.in/mgo.v2/bson"
)

type Share struct {
	*revel.Controller
}

type Showshare struct {
	models.Share
	Name string
}
func (c Share) Minna() revel.Result {

	db := c.RenderArgs["mgo"].(models.MyMgo).DB(c.RenderArgs["db"].(string))
	var els []models.Share
	db.C(models.SH).Find(nil).Sort("-_id").All(&els)
	var eels []Showshare
	eels = make([]Showshare,0)
	for _,v := range els {
		eels = append(eels, Showshare{Share:v,Name:models.GetName(db, v.ShijuanId)})
	}
	c.RenderArgs["list1"] = els
	c.RenderArgs["list"] = eels
	return c.Render()
}
func (c Share) To(id string) revel.Result {
	db := c.RenderArgs["mgo"].(models.MyMgo).DB(c.RenderArgs["db"].(string))
	var els models.Share
	sid := bson.NewObjectId()
	db.C(models.SH).Find(bson.M{"_id":bson.ObjectIdHex(id)}).One(&els)
	db.C(models.SH).Insert(&models.Share{
		Id_:sid,
		ShijuanId:id,
	})
	return c.Redirect("/share/minna")
}
