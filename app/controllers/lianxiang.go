package controllers

import (
	"github.com/revel/revel"
	"github.com/zhangsong/selfup/app/models"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type Lianxiang struct {
	*revel.Controller
}

func (c Lianxiang) Index(id string) revel.Result {
	c.RenderArgs["lid"] = id
	return c.Render()
}
func (c Lianxiang) Adddo(lid, content string) revel.Result {

	if c.Request.Method != "POST" {
		return c.Redirect("/")
	}
	db := c.RenderArgs["mgo"].(models.MyMgo).DB(c.RenderArgs["db"].(string))

	id := bson.NewObjectId()
	db.C(models.LX).Insert(&models.Lianxiang{
		Id_:id,
		Content:content,
		Oid:lid,
	})
	return c.Redirect("/")
}
func (c Lianxiang) Show(id string) revel.Result {

	db := c.RenderArgs["mgo"].(models.MyMgo).DB(c.RenderArgs["db"].(string))

	var els []models.Lianxiang
	db.C(models.LX).Find(bson.M{"oid":id}).All(&els)
	//db.C(models.LX).Find(nil).All(&els)
	var elss []string
	elss = make([]string, 0)
	for _, v := range els {
		elss = append(elss, v.Content)
	}
	fmt.Println(id)
	return c.RenderJson(elss)
}
