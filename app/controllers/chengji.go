package controllers

import (
	"github.com/revel/revel"
	"github.com/zhangsong/selfup/app/models"
)

type Chengji struct {
	*revel.Controller
}


func (c Chengji) Index() revel.Result {

	db := c.RenderArgs["mgo"].(models.MyMgo).DB(c.RenderArgs["db"].(string))
	var els []models.Chengji
	db.C(models.CJ).Find(nil).Sort("-_id").All(&els)
	c.RenderArgs["list"] = els
	return c.Render()
}
