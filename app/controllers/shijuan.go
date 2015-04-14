package controllers

import (
	"github.com/revel/revel"
	"github.com/zhangsong/selfup/app/models"
	"fmt"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
)

type Shijuan struct {
	*revel.Controller
}


func (c Shijuan) Add() revel.Result {

	db := c.RenderArgs["mgo"].(models.MyMgo).DB(c.RenderArgs["db"].(string))
	var els []models.Element
	db.C(models.EL).Find(nil).All(&els)
	c.RenderArgs["list"] = els
	return c.Render()
}
func (c Shijuan) My() revel.Result {

	db := c.RenderArgs["mgo"].(models.MyMgo).DB(c.RenderArgs["db"].(string))
	var els []models.Shijuan
	db.C(models.SJ).Find(nil).All(&els)
	c.RenderArgs["list"] = els
	return c.Render()
}
func (c Shijuan) Adddo(shijuan []string) revel.Result {

	if c.Request.Method != "POST" {
		return c.Redirect("/")
	}

	/*
	db := c.RenderArgs["mgo"].(models.MyMgo).DB(c.RenderArgs["db"].(string))
	var els []models.Element
	db.C(models.EL).Find(nil).All(&els)
	c.RenderArgs["list"] = els
	*/
	if mb, err := json.Marshal(shijuan); err == nil {
		db := c.RenderArgs["mgo"].(models.MyMgo).DB(c.RenderArgs["db"].(string))
		id := bson.NewObjectId()
		err = db.C(models.SJ).Insert(&models.Shijuan{
			Id_:id,
			Content:string(mb),
		})
		if err != nil {
			return c.Redirect("/")
		}
		fmt.Println(err)

	}
	return c.Redirect("/shijuan/my")
}
func (c Shijuan) Ask() revel.Result {

	db := c.RenderArgs["mgo"].(models.MyMgo).DB(c.RenderArgs["db"].(string))
	var els models.Shijuan
	db.C(models.SJ).Find(bson.M{"_id":bson.ObjectIdHex(c.Params.Get("id"))}).One(&els)

	var sj []string
	if err := json.Unmarshal([]byte(els.Content), &sj); err != nil {
		return c.Redirect("/")
	}
	
	var ee []models.Element

	ee = make([]models.Element, 0)

	for _,id := range sj {
		var e models.Element
		db.C(models.EL).Find(bson.M{"_id":bson.ObjectIdHex(id)}).One(&e)
		ee = append(ee, e)
	}
	c.RenderArgs["shijuan"] = els
	c.RenderArgs["ask"] = ee
	return c.Render()
}
func (c Shijuan) Ceyan() revel.Result {
	if c.Request.Method != "POST" {
		return c.Redirect("/")
	}

	db := c.RenderArgs["mgo"].(models.MyMgo).DB(c.RenderArgs["db"].(string))
	var els models.Shijuan
	if err := db.C(models.SJ).Find(bson.M{"_id":bson.ObjectIdHex(c.Params.Get("id"))}).One(&els); err != nil {
		return c.Redirect("/")
	}
//成绩判定
id := bson.NewObjectId()
err := db.C(models.CJ).Insert(&models.Chengji{
		Id_:id,
		Count:0,
		ShijuanId:c.Params.Get("id"),
	})
	if err != nil {
		return c.Redirect("/")
	}
	return c.Redirect("/chengji/index")
}
