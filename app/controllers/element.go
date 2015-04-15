package controllers

import (
	"github.com/revel/revel"
	//"github.com/dchest/captcha"
	"strings"
	"fmt"
	"io"
	"encoding/json"
	"crypto/sha1"
	//"errors"
	"gopkg.in/mgo.v2/bson"
	"github.com/zhangsong/selfup/app/models"
)

type Element struct {
	*revel.Controller
}


func (c Element) Add() revel.Result {
	/*var members []models.User
	db := c.RenderArgs["mgo"].(models.MyMgo).DB(c.RenderArgs["db"].(string))
	db.C(models.USERS).Find(nil).Sort("-username").Limit(40).All(&members)
	membersCount, _ := db.C(models.USERS).Find(nil).Count()
	*/
	return c.Render()
}
func (c Element) Adddo() revel.Result {
	if c.Request.Method != "POST" {
		return c.Redirect("/")
	}

	el := strings.TrimSpace(c.Params.Get("el"))
	c.Validation.Required(el)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect("/element/add")
	}

	daan := models.GetAanda(el)
	var dd string
	if bb,err :=json.Marshal(daan); err != nil || len(bb)==0 {
		return c.Redirect("/")
	} else {
		dd = string(bb)
	}

	h := sha1.New()
	io.WriteString(h, el)
	db := c.RenderArgs["mgo"].(models.MyMgo).DB(c.RenderArgs["db"].(string))
	id := bson.NewObjectId()
	err := db.C(models.EL).Insert(&models.Element{
		Id_:id,
		Content:el,
		Hash:fmt.Sprintf("%x",h.Sum(nil)),
		Aanda:dd,
	})

	if err != nil {
		return c.Redirect("/element/add")
	}
	return c.Redirect("/element/add")
}
func (c Element) My() revel.Result {
	db := c.RenderArgs["mgo"].(models.MyMgo).DB(c.RenderArgs["db"].(string))
	var els []models.Element
	db.C(models.EL).Find(nil).All(&els)
	c.RenderArgs["list"] = els
	return c.Render()
}
func (c Element) Delete() revel.Result {
	db := c.RenderArgs["mgo"].(models.MyMgo).DB(c.RenderArgs["db"].(string))
	err := db.C(models.EL).Remove(bson.M{"_id":bson.ObjectIdHex(c.Params.Get("id"))})
	if err != nil {
		return c.RenderError(err)
	}
	return c.Redirect("/element/my")
}
