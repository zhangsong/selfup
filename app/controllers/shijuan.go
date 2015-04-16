package controllers

import (
	"github.com/revel/revel"
	"github.com/zhangsong/selfup/app/models"
	"fmt"
	"strings"
	"encoding/json"
	"html/template"
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
	db.C(models.SJ).Find(nil).Sort("-_id").All(&els)
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
		var name string
		if c.Params.Get("makename") == "name" && len(strings.TrimSpace(c.Params.Get("shijuanname")))>0 {
			name = strings.TrimSpace(c.Params.Get("shijuanname"))
		}
		db := c.RenderArgs["mgo"].(models.MyMgo).DB(c.RenderArgs["db"].(string))
		id := bson.NewObjectId()
		err = db.C(models.SJ).Insert(&models.Shijuan{
			Id_:id,
			Content:string(mb),
			Name:name,
		})
		if err != nil {
			return c.Redirect("/")
		}
		fmt.Println(err)

	}
	return c.Redirect("/shijuan/my")
}

type Myhtml struct {
	Element models.Element
	Htmlcontent template.HTML
}
func (c Shijuan) Ask() revel.Result {

	db := c.RenderArgs["mgo"].(models.MyMgo).DB(c.RenderArgs["db"].(string))
	var els models.Shijuan
	db.C(models.SJ).Find(bson.M{"_id":bson.ObjectIdHex(c.Params.Get("id"))}).One(&els)

	var sj []string
	if err := json.Unmarshal([]byte(els.Content), &sj); err != nil {
		return c.Redirect("/")
	}
	
	var ee []Myhtml

	ee = make([]Myhtml, 0)

	for _,id := range sj {
		var e models.Element
		db.C(models.EL).Find(bson.M{"_id":bson.ObjectIdHex(id)}).One(&e)
		ee = append(ee, Myhtml{Element:e, Htmlcontent:template.HTML(models.GenHtml(e.Content, e.Id_.Hex()))})
	}
	fmt.Println(ee)
	c.RenderArgs["shijuan"] = els
	c.RenderArgs["ask"] = ee
	return c.Render()
}
func (c Shijuan) Ceyan(sub []string) revel.Result {
	if c.Request.Method != "POST" {
		return c.Redirect("/")
	}

	db := c.RenderArgs["mgo"].(models.MyMgo).DB(c.RenderArgs["db"].(string))
	var els models.Shijuan
	if err := db.C(models.SJ).Find(bson.M{"_id":bson.ObjectIdHex(c.Params.Get("id"))}).One(&els); err != nil {
		return c.Redirect("/")
	}

	var cc = 0
	for _, v := range sub {
		var jd []string
		c.Params.Bind(&jd, "vv"+v)
		if _, y := models.Juge(db, v, jd);y {
			cc++
		}
	}
//成绩判定
id := bson.NewObjectId()
err := db.C(models.CJ).Insert(&models.Chengji{
		Id_:id,
		Count:cc,
		ShijuanId:c.Params.Get("id"),
	})
	if err != nil {
		return c.Redirect("/")
	}
	return c.Redirect("/chengji/index")
}
